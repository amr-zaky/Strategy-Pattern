package strategy

type Strategy interface {
	DoOperation(number1 int, number2 int) int
}
