package commands

import (
	"fmt"
	"os/exec"

	"github.com/artarts36/quicktool/internal/presentation/interaction"
)

type Gpg struct {
}

func NewGpg() *Gpg {
	return &Gpg{}
}

func (c *Gpg) Definition() *interaction.Definition {
	return &interaction.Definition{
		Name:        "gpg",
		Description: "gpg",
		Args: []*interaction.DefinitionArg{
			{
				Name: "action",
				ValuesEnum: []string{
					"pk",
				},
				Required: true,
			},
			{
				Name:     "key-id",
				Required: true,
			},
		},
	}
}

func (c *Gpg) Execute(ctx *interaction.Context, env *interaction.Env) error {
	action := env.Input.Argument("action")
	keyID := env.Input.Argument("key-id")
	if action == "pk" {
		cmd := exec.CommandContext(ctx.Context, "gpg", "--export-secret-key", "-a", keyID)
		res, err := cmd.Output()
		if err != nil {
			return err
		}

		fmt.Printf("%s", res)

		return nil
	}

	return fmt.Errorf("action %s unsupported", action)
}
