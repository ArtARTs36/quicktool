package commands

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/artarts36/quicktool/internal/presentation/interaction"
	"github.com/artarts36/quicktool/internal/shared"
)

type Mkdirs struct {
}

func NewMkdirsCommand() *Mkdirs {
	return &Mkdirs{}
}

func (c *Mkdirs) Definition() *interaction.Definition {
	return &interaction.Definition{
		Name:        "mkdirs",
		Description: "Make directories",
		Args: []*interaction.DefinitionArg{
			{
				Name:        "range",
				Description: "Range, example \"0-20\"",
				Required:    true,
			},
			{
				Name:        "fileMask",
				Description: "Mask to folder name, example \"folder{number}\"",
				Required:    false,
			},
		},
	}
}

func (c *Mkdirs) Execute(_ *interaction.Context, env *interaction.Env) error {
	rangeStr := env.Input.Argument("range")
	rangeVal, err := shared.RangeFromString(rangeStr)
	if err != nil {
		return fmt.Errorf("invalid range: %s", err.Error())
	}

	fmask := env.Input.Argument("fileMask")
	if fmask == "" {
		for i := rangeVal.From; i < rangeVal.To; i++ {
			name := fmt.Sprintf("%d", i)

			err := os.Mkdir(name, 0700)
			if err != nil {
				return err
			}
		}

		return nil
	}

	for i := rangeVal.From; i < rangeVal.To; i++ {
		name := strings.ReplaceAll(fmask, "{number}", strconv.Itoa(i))

		err := os.MkdirAll(name, 0700)
		if err != nil {
			return err
		}
	}

	return nil
}
