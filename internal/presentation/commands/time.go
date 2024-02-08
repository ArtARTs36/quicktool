package commands

import (
	"fmt"
	"time"

	"github.com/artarts36/quicktool/internal/presentation/interaction"
)

type Timestamp struct {
}

func NewTimestampCommand() *Timestamp {
	return &Timestamp{}
}

func (c *Timestamp) Definition() *interaction.Definition {
	return &interaction.Definition{
		Name:        "time",
		Description: "Get time",
	}
}

func (c *Timestamp) Execute(_ *interaction.Context, env *interaction.Env) error {
	t := time.Now()

	env.PrintTable(
		[]string{
			"type",
			"value",
		},
		[][]string{
			{
				"Y-m-d H:i:s",
				t.Format(time.DateTime),
			},
			{
				"Unix",
				fmt.Sprintf("%d", t.Unix()),
			},
		},
	)

	return nil
}
