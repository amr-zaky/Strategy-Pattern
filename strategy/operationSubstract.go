package strategy

type OperationSubstract struct {
}

func (s *OperationSubstract) DoOperation(number1 int, number2 int) int {
	return number1 - number2
}
