package commands

import (
	"fmt"

	"github.com/artarts36/quicktool/internal/shared"

	"github.com/artarts36/quicktool/internal/presentation/interaction"
)

type JSONPretty struct {
}

func NewJSONPrettyCommand() *JSONPretty {
	return &JSONPretty{}
}

func (c *JSONPretty) Definition() *interaction.Definition {
	return &interaction.Definition{
		Name:        "jsonpretty",
		Description: "Pretty json",
		Args: []*interaction.DefinitionArg{
			{
				Name:        "value",
				Description: "source value",
				Required:    true,
			},
		},
	}
}

func (c *JSONPretty) Execute(_ *interaction.Context, env *interaction.Env) error {
	source := env.Input.Argument("value")

	json, err := shared.PrettyJSON(source)
	if err != nil {
		return fmt.Errorf("json invalid: %s", err.Error())
	}

	env.PrintText(json)

	return nil
}
