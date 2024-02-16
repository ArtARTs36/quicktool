package shared

import (
	"fmt"
	"strconv"
)

const deltaTokensCount = 2

type Delta struct {
	Value     int
	Operation DeltaOperation
}

func NewDeltaFromString(str string) (*Delta, error) {
	if len(str) < deltaTokensCount {
		return nil, fmt.Errorf("str must be %d or greather length", deltaTokensCount)
	}

	op, opExists := deltaOperationsMap[string(str[0])]
	if !opExists {
		return nil, fmt.Errorf("operation %q unsupported", str[0])
	}

	val, err := strconv.Atoi(str[1:])
	if err != nil {
		return nil, fmt.Errorf("invalid value: %s", err)
	}

	return &Delta{
		Value:     val,
		Operation: op,
	}, nil
}

func DeltaZero() *Delta {
	return &Delta{
		Value:     0,
		Operation: DeltaOperationAdd,
	}
}

func (d *Delta) ExecuteOperation(value int) int {
	return d.Operation.Execute(value, d.Value)
}
