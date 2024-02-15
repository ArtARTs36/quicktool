package shared

import (
	"fmt"
	"strconv"
	"strings"
)

const rangeTokensCount = 2

type Range struct {
	From int
	To   int
}

func RangeFromString(value string) (*Range, error) {
	tokens := strings.Split(value, "-")
	if len(tokens) != rangeTokensCount {
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

func (r *Range) Len() int {
	return max(r.From, r.To) - min(r.From, r.To)
}

func (r *Range) IsDescending() bool {
	return r.To < r.From
}

func (r *Range) Iterate(fn func(int) error) error {
	if r.IsDescending() {
		for i := r.From; i >= r.To; i-- {
			err := fn(i)
			if err != nil {
				return err
			}
		}

		return nil
	}

	for i := r.From; i <= r.To; i++ {
		err := fn(i)
		if err != nil {
			return err
		}
	}

	return nil
}
