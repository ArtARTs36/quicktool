package commands

import (
	"fmt"
	"strings"

	"github.com/artarts36/quicktool/internal/presentation/interaction"
	"github.com/artarts36/quicktool/internal/shared"
)

type Range struct {
}

func NewRange() *Range {
	return &Range{}
}

func (r *Range) Definition() *interaction.Definition {
	return &interaction.Definition{
		Name:        "range",
		Description: "generate range",
		Args: []*interaction.DefinitionArg{
			{
				Name:        "length",
				Required:    true,
				Description: "Range, example: 1-10",
			},
			{
				Name:        "delta",
				Required:    false,
				Description: "delta value, example: +2",
			},
		},
	}
}

func (r *Range) Execute(_ *interaction.Context, env *interaction.Env) error {
	length, err := shared.RangeFromString(env.Input.Argument("length"))
	if err != nil {
		return fmt.Errorf("unable to parse range string: %s", err)
	}

	delta := shared.DeltaZero()

	deltaVal := env.Input.Argument("delta")
	if deltaVal != "" {
		delta, err = shared.NewDeltaFromString(deltaVal)
		if err != nil {
			return fmt.Errorf("unable to parse delta string: %s", err)
		}
	}

	numbers := make([]string, 0, length.Len())

	err = length.Iterate(func(i int) error {
		numbers = append(numbers, fmt.Sprintf("%d", delta.ExecuteOperation(i)))
		return nil
	})
	if err != nil {
		return err
	}

	env.PrintText(fmt.Sprintf("[%s]", strings.Join(numbers, ",")))

	return nil
}
