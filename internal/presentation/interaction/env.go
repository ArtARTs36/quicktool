package interaction

import (
	"bufio"
	"bytes"
	"os"

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

func (e *Env) ReadMultiline() ([]byte, error) {
	dropCR := func(data []byte) []byte {
		if len(data) > 0 && data[len(data)-1] == '\r' {
			return data[0 : len(data)-1]
		}
		return data
	}

	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		if i := bytes.Index(data, []byte{'\n', '\n'}); i >= 0 {
			return i + 2, dropCR(data[0:i]), nil
		}

		// If we're at EOF, we have a final, non-terminated line. Return it.
		if atEOF {
			return len(data), dropCR(data), nil
		}
		// Request more data.
		return 0, nil, nil
	})

	if scanner.Scan() {
		b := scanner.Bytes()

		return b, scanner.Err()
	}

	return []byte{}, scanner.Err()
}
