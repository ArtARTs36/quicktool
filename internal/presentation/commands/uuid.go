package commands

import (
	"github.com/artarts36/quicktool/internal/application"
	"github.com/artarts36/quicktool/internal/presentation/interaction"
)

type UUID struct {
}

func NewUUIDCommand() *UUID {
	return &UUID{}
}

func (c *UUID) Definition() *interaction.Definition {
	return &interaction.Definition{
		Name:        "uuid",
		Description: "Generate UUID",
	}
}

func (c *UUID) Execute(_ *interaction.Context, env *interaction.Env) error {
	env.PrintTable(
		[]string{
			"version", "value",
		},
		[][]string{
			{
				"v4",
				application.GenerateUUIDV4().String(),
			},
			{
				"v6",
				application.GenerateUUIDV6().String(),
			},
			{
				"v7",
				application.GenerateUUIDV7().String(),
			},
		},
	)

	return nil
}
