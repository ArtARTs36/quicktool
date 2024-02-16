package main

import (
	"log"
	"os"

	"github.com/artarts36/quicktool/internal/presentation/commands"
	"github.com/artarts36/quicktool/internal/presentation/interaction"
	"github.com/tyler-sommer/stick"
)

func main() {
	twig := stick.New(stick.NewFilesystemLoader("./"))

	actions := []struct {
		templateFilepath string
		outputFilepath   string
		action           func() map[string]stick.Value
	}{
		{
			templateFilepath: "./docs/readme_template.md.twig",
			outputFilepath:   "README.md",
			action: func() map[string]stick.Value {
				cmds := commands.CreateList(nil)
				commandDefinitions := make([]*interaction.Definition, 0, len(cmds))
				for _, cmd := range cmds {
					commandDefinitions = append(commandDefinitions, cmd.Definition())
				}

				return map[string]stick.Value{
					"sh": commandDefinitions,
				}
			},
		},
	}

	for _, action := range actions {
		outFile, err := os.Create(action.outputFilepath)
		if err != nil {
			log.Fatal(err)
		}

		err = twig.Execute(action.templateFilepath, outFile, action.action())
		if err != nil {
			log.Fatal(err)
		}
	}
}
