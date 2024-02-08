package interaction

import (
	go_console "github.com/DrSmithFr/go-console"
	"github.com/DrSmithFr/go-console/table"
)

type Env struct {
	go_console.Script
}

func (e *Env) PrintTable(headers []string, rows [][]string) {
	render := table.
		NewRender(e.Output).
		SetStyleFromName("box")

	render.SetContent(
		table.
			NewTable().
			SetHeaders(
				table.MakeDataFromStrings([][]string{headers}),
			).
			SetRowsFromString(rows),
	)

	render.Render()
}
