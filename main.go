package main

import (
	"fmt"
	"test2/strategy"
)

func main() {

	addOperation := strategy.OperationAdd{}
	subtractOperation := strategy.OperationSubstract{}
	multiplyOperation := strategy.OperationMultiply{}

	context := strategy.Context{
		Strategy: &addOperation,
	}
	fmt.Println("do operation Add:", context.ExecuteStrategy(3, 4))
	context = strategy.Context{
		Strategy: &subtractOperation,
	}
	fmt.Println("do operation Subtract:", context.ExecuteStrategy(3, 4))
	context = strategy.Context{
		Strategy: &multiplyOperation,
	}

	fmt.Println("do operation multiply:", context.ExecuteStrategy(3, 4))
}
