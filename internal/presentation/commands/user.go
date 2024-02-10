package commands

import "github.com/artarts36/quicktool/internal/presentation/interaction"

type User struct {
}

func NewUser() *User {
	return &User{}
}

func (c *User) Definition() *interaction.Definition {
	return &interaction.Definition{
		Name:        "user",
		Description: "Get user data",
	}
}

func (c *User) Execute(ctx *interaction.Context, env *interaction.Env) error {
	env.PrintTable(
		[]string{
			"Field",
			"Value",
		},
		[][]string{
			{
				"Name",
				ctx.User.Name,
			},
			{
				"Home directory",
				ctx.User.HomeDir,
			},
		},
	)

	return nil
}
