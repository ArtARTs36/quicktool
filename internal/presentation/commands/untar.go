package commands

import (
	"fmt"

	"github.com/artarts36/quicktool/internal/infrastructure/sh"
	"github.com/artarts36/quicktool/internal/presentation/interaction"
)

type Untar struct {
	tar *sh.Tar
}

func NewUntar(tar *sh.Tar) *Untar {
	return &Untar{
		tar: tar,
	}
}

func (t *Untar) Definition() *interaction.Definition {
	return &interaction.Definition{
		Name:        "untar",
		Description: "Extract tar archive",
		Args: []*interaction.DefinitionArg{
			{
				Name:        "archive-path",
				Description: "path to archive",
				Required:    true,
			},
		},
	}
}

func (t *Untar) Execute(ctx *interaction.Context, env *interaction.Env) error {
	res, err := t.tar.Extract(ctx.Context, &sh.TarExtractParams{
		ArchivePath: env.Input.Argument("archive-path"),
	})
	if err != nil {
		if len(res) > 0 {
			env.PrintText(fmt.Sprintf("Error: %s", res))
		}

		return err
	}

	env.PrintText(fmt.Sprintf("Result: %s", res))

	return nil
}
