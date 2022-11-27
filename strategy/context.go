package strategy

type Context struct {
	Strategy Strategy
}

func (c Context) ExecuteStrategy(number1 int, number2 int) int {
	return c.Strategy.DoOperation(number1, number2)
}
