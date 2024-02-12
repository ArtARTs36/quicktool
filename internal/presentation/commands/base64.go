package commands

import (
	"encoding/base64"
	"fmt"

	"github.com/artarts36/quicktool/internal/presentation/interaction"
)

type Base64 struct {
}

func NewBase64() *Base64 {
	return &Base64{}
}

func (c *Base64) Definition() *interaction.Definition {
	return &interaction.Definition{
		Name:        "base64",
		Description: "Base64 encode/decode",
		Args: []*interaction.DefinitionArg{
			{
				Name:        "action",
				Description: "action",
				ValuesEnum: []string{
					"encode",
					"decode",
				},
				Required: true,
			},
			{
				Name:        "value",
				Description: "value to encode or encode",
				Required:    true,
			},
		},
	}
}

func (c *Base64) Execute(ctx *interaction.Context, env *interaction.Env) error {
	action := env.Input.Argument("action")
	value := env.Input.Argument("value")

	if action == "decode" {
		decoded, err := base64.StdEncoding.DecodeString(value)
		if err != nil {
			return fmt.Errorf("unable decode base64: %s", err)
		}

		env.PrintText(fmt.Sprintf("Decoded value: %s", decoded))

		return nil
	}

	env.PrintText(fmt.Sprintf("Encoded value: %s", base64.StdEncoding.EncodeToString([]byte(value))))

	return nil
}
