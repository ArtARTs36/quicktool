package shared

import (
	"fmt"
	"strconv"
	"strings"
)

type Range struct {
	From int
	To   int
}

func RangeFromString(value string) (*Range, error) {
	tokens := strings.Split(value, "-")
	if len(tokens) != 2 {
		return nil, fmt.Errorf("must be 2 values")
	}

	from, err := strconv.Atoi(tokens[0])
	if err != nil {
		return nil, fmt.Errorf("\"from\" must be integer")
	}

	to, err := strconv.Atoi(tokens[1])
	if err != nil {
		return nil, fmt.Errorf("\"to\" must be integer")
	}

	return &Range{
		From: from,
		To:   to,
	}, nil
}
