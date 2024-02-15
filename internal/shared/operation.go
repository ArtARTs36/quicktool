package shared

type DeltaOperation string

const (
	DeltaOperationAdd DeltaOperation = "+"
	DeltaOperationSub DeltaOperation = "-"
	DeltaOperationMul DeltaOperation = "*"
	DeltaOperationDiv DeltaOperation = "/"
)

var deltaOperationsMap = map[string]DeltaOperation{
	"+": DeltaOperationAdd,
	"-": DeltaOperationSub,
	"*": DeltaOperationMul,
	"/": DeltaOperationDiv,
}

func (d DeltaOperation) Execute(first, second int) int {
	switch d {
	case DeltaOperationDiv:
		return first / second
	case DeltaOperationMul:
		return first * second
	case DeltaOperationSub:
		return first - second
	case DeltaOperationAdd:
		return first + second
	default:
		return first + second
	}
}
