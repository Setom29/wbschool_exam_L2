package pattern

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

// https://golangbyexample.com/strategy-design-pattern-golang/

import "fmt"

type Operator interface {
	Apply(int, int) int
}

type Operation struct {
	Operator Operator
}

func (o *Operation) Operate(leftValue, rightValue int) int {
	return o.Operator.Apply(leftValue, rightValue)
}

// addition

type Addition struct{}

func (Addition) Apply(lval, rval int) int {
	return lval + rval
}

// multiplication

type Multiplication struct{}

func (Multiplication) Apply(lval, rval int) int {
	return lval * rval
}

func main() {
	add := Operation{Addition{}}
	fmt.Println(add.Operate(3, 5)) // 8

	mult := Operation{Multiplication{}}
	fmt.Println(mult.Operate(3, 5))
}
