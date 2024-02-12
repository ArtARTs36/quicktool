package commands

import (
	"encoding/base64"
	"fmt"

	"github.com/artarts36/quicktool/internal/infrastructure/filesystem"
	"github.com/artarts36/quicktool/internal/presentation/interaction"
)

const (
	base64ActionEncode = "encode"
	base64ActionDecode = "decode"

	base64OutputStdout = "stdout"
	base64OutputFile   = "file"
)

type Base64 struct {
	fs filesystem.FileSystem
}

func NewBase64(fs filesystem.FileSystem) *Base64 {
	return &Base64{
		fs: fs,
	}
}

func (c *Base64) Definition() *interaction.Definition {
	return &interaction.Definition{
		Name:        "base64",
		Description: "Base64 encode/decode",
		Aliases: []string{
			"b64",
		},
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
		Opts: []*interaction.DefinitionOpt{
			{
				Name:        "output",
				ShortName:   "O",
				Description: "path for output file",
			},
		},
	}
}

func (c *Base64) Execute(_ *interaction.Context, env *interaction.Env) error {
	action := env.Input.Argument("action")
	if action == "" {
		action = base64ActionDecode
	}

	outputType := base64OutputStdout
	output := env.Input.Option("output")
	if output != "" {
		outputType = base64OutputFile
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

	actionFn := c.decode
	if action == base64ActionEncode {
		actionFn = c.encode
	}

	err = actionFn(env, source, outputType, output)
	if err != nil {
		return fmt.Errorf("unable to %s: %s", action, err)
	}

	return nil
}

func (c *Base64) encode(env *interaction.Env, source []byte, outputType, output string) error {
	encoded := base64.StdEncoding.EncodeToString(source)

	if outputType == base64OutputStdout {
		env.PrintText(fmt.Sprintf("Encoded value: %s", encoded))

		return nil
	}

	err := c.fs.Save(output, []byte(encoded))
	if err != nil {
		return fmt.Errorf("unable to save file: %s", err)
	}

	env.PrintText(fmt.Sprintf("Encoded value stored in %s", output))

	return nil
}

func (c *Base64) decode(env *interaction.Env, source []byte, outputType, output string) error {
	decoded, err := base64.StdEncoding.DecodeString(string(source))
	if err != nil {
		return fmt.Errorf("decoding failed: %s", err)
	}

	if outputType == base64OutputStdout {
		env.PrintText(fmt.Sprintf("Decoded value: %s", decoded))

		return nil
	}

	err = c.fs.Save(output, decoded)
	if err != nil {
		return fmt.Errorf("unable to save file: %s", err)
	}

	env.PrintText(fmt.Sprintf("Decoded value stored in %s", output))

	return nil
}
