package commands

import (
	"fmt"

	"github.com/artarts36/quicktool/internal/shared"

	"github.com/artarts36/quicktool/internal/presentation/interaction"
)

type MD5 struct {
}

func NewMD5() *MD5 {
	return &MD5{}
}

func (c *MD5) Definition() *interaction.Definition {
	return &interaction.Definition{
		Name:        "md5",
		Description: "Encode value to md5",
		Args: []*interaction.DefinitionArg{
			{
				Name:        "value",
				Description: "source value",
				Required:    true,
			},
		},
	}
}

func (c *MD5) Execute(_ *interaction.Context, env *interaction.Env) error {
	source := env.Input.Argument("value")

	hash := shared.HashMD5(source)

	env.PrintText(fmt.Sprintf("md5: %s", hash))

	return nil
}
