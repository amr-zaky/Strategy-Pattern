package strategy

type OperationMultiply struct {
}

func (m *OperationMultiply) DoOperation(number1 int, number2 int) int {
	return number1 * number2
}
