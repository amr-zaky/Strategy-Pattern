package strategy

type OperationAdd struct {
}

func (a *OperationAdd) DoOperation(number1 int, number2 int) int {
	return number1 + number2
}
