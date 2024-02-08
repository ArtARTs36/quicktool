package commands

import (
	"fmt"
	"strconv"

	"github.com/sethvargo/go-password/password"

	"github.com/artarts36/quicktool/internal/presentation/interaction"
)

const (
	defaultPasswordLength = 12
	symbolsPerDigit       = 4
)

type Password struct {
}

func NewPasswordCommand() *Password {
	return &Password{}
}

func (p *Password) Definition() *interaction.Definition {
	return &interaction.Definition{
		Name: "password",
		Aliases: []string{
			"pswd",
		},
		Description: "Generate password",
		Args: []*interaction.DefinitionArg{
			{
				Name:        "length",
				Description: "length of password",
				Required:    false,
			},
		},
	}
}

func (p *Password) Execute(_ *interaction.Context, env *interaction.Env) error {
	length, err := p.length(env)
	if err != nil {
		return err
	}

	numDigits := length / symbolsPerDigit
	numSymbols := length - numDigits

	pswd, err := password.Generate(length, numDigits, numSymbols, false, true)
	if err != nil {
		return fmt.Errorf("unable to generate password: %s", err)
	}

	env.PrintInfoSubject("password", pswd)

	return nil
}

func (p *Password) length(env *interaction.Env) (int, error) {
	length := env.Input.Argument("length")
	if length == "" {
		return defaultPasswordLength, nil
	}

	lengthNum, err := strconv.Atoi(length)
	if err != nil {
		return 0, fmt.Errorf("invalid password length: %s", err.Error())
	}

	return lengthNum, nil
}
