package application

import (
	"bytes"
	"encoding/json"
)

func PrettyJSON(source string) (string, error) {
	var buf bytes.Buffer

	if err := json.Indent(&buf, []byte(source), "", "    "); err != nil {
		return "", err
	}

	return buf.String(), nil
}
