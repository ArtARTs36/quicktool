package interaction

import (
	"fmt"
	"strings"

	"github.com/DrSmithFr/go-console/input/option"

	goconsole "github.com/DrSmithFr/go-console"
	"github.com/DrSmithFr/go-console/input/argument"
)

type Console struct {
	cmd *goconsole.Command

	currentContext *Context
}

type AppInfo struct {
	Name      string
	Version   string
	BuildDate string
}

func NewConsole(appInfo *AppInfo, commands []Command) *Console {
	console := &Console{}
	console.cmd = &goconsole.Command{
		Scripts: console.buildGoConsoleScripts(commands),
		BuildInfo: &goconsole.BuildInfo{
			Name:      appInfo.Name,
			Version:   appInfo.Version,
			BuildFlag: appInfo.BuildDate,
		},
	}

	return console
}

func (c *Console) buildGoConsoleScripts(commands []Command) []*goconsole.Script {
	scripts := make([]*goconsole.Script, 0)

	for _, command := range commands {
		command := command
		def := command.Definition()

		args := make([]goconsole.Argument, 0, len(def.Args))
		opts := make([]goconsole.Option, 0, len(def.Opts))

		for _, arg := range def.Args {
			argVal := argument.Optional

			if arg.Required {
				argVal = argument.Required
			}

			description := arg.Description
			if len(arg.ValuesEnum) > 0 {
				description = fmt.Sprintf(
					"%s: value of [%s]",
					description,
					strings.Join(arg.ValuesEnum, ", "),
				)
			}

			args = append(args, goconsole.Argument{
				Name:        arg.Name,
				Description: description,
				Value:       argVal,
			})
		}

		for _, opt := range def.Opts {
			opts = append(opts, goconsole.Option{
				Name:        opt.Name,
				Shortcut:    opt.ShortName,
				Description: opt.Description,
				Value:       option.Optional,
			})
		}

		script := &goconsole.Script{
			Name:        def.Name,
			Description: def.Description,
			Runner: func(script *goconsole.Script) goconsole.ExitCode {
				return c.runCommand(command, script)
			},
			Arguments: args,
			Options:   opts,
		}

		scripts = append(scripts, script)

		for _, alias := range def.Aliases {
			scripts = append(scripts, &goconsole.Script{
				Name:        alias,
				Description: fmt.Sprintf("%s (alias to %s)", def.Description, def.Name),
				Runner: func(script *goconsole.Script) goconsole.ExitCode {
					return c.runCommand(command, script)
				},
				Arguments: args,
			})
		}
	}

	return scripts
}

func (c *Console) Run(ctx *Context) {
	c.currentContext = ctx

	c.cmd.Run()
}

func (c *Console) runCommand(cmd Command, script *goconsole.Script) goconsole.ExitCode {
	err := cmd.Execute(c.currentContext, &Env{
		Script: *script,
	})
	if err != nil {
		script.PrintError(fmt.Sprintf("command failed: %s", err.Error()))

		return goconsole.ExitError
	}
	return goconsole.ExitSuccess
}
