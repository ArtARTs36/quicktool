package commands

import (
	"encoding/base64"
	"fmt"

	"github.com/artarts36/quicktool/internal/presentation/interaction"
)

const (
	base64ActionEncode = "encode"
	base64ActionDecode = "decode"
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
				Required:    false,
				ValuesEnum: []string{
					base64ActionEncode,
					base64ActionDecode,
				},
			},
		},
	}
}

func (c *Base64) Execute(_ *interaction.Context, env *interaction.Env) error {
	action := env.Input.Argument("action")
	if action == "" {
		action = base64ActionDecode
	}

	if action == base64ActionEncode {
		env.PrintText("Enter data to encode: ")
	} else {
		env.PrintText("Enter data to decode: ")
	}

	source, err := env.ReadMultiline()
	if err != nil {
		return fmt.Errorf("unable to read input source: %s", err.Error())
	}

	if action == base64ActionEncode {
		env.PrintText(fmt.Sprintf("Encoded value: %s", base64.StdEncoding.EncodeToString(source)))
	} else {
		res, err := base64.StdEncoding.DecodeString(string(source))
		if err != nil {
			return fmt.Errorf("decoding failed: %s", err)
		}

		env.PrintText(fmt.Sprintf("Decoded value: %s", res))
	}

	return nil
}
