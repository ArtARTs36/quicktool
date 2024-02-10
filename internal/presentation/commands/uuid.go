package commands

import (
	"github.com/artarts36/quicktool/internal/presentation/interaction"
	"github.com/artarts36/quicktool/internal/shared"
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
				shared.GenerateUUIDV4().String(),
			},
			{
				"v6",
				shared.GenerateUUIDV6().String(),
			},
			{
				"v7",
				shared.GenerateUUIDV7().String(),
			},
		},
	)

	return nil
}
