package main

import (
	"fmt"
	"reflect"
)

// Паттерн Strategy относится к поведенческим паттернам
// Паттерн применяется в случаях, когда алгоритм поведения необходимо определить в процессе выполнения программы в зависимости от контекста

// Пример - определение формата документа, алгоритма кодирования и т.п.
type operator interface {
	Apply(int, int) int
}
type Operation struct {
	name string
	op   operator
}

func NewOperation(op operator) *Operation {
	return &Operation{
		name: reflect.ValueOf(op).String(),
		op:   op,
	}
}
func (o *Operation) GetName() string {
	return o.name
}
func (o *Operation) Operate(leftValue, rightValue int) int {
	return o.op.Apply(leftValue, rightValue)
}

type Addition struct{}

func (Addition) Apply(lval, rval int) int {
	return lval + rval
}

type Substraction struct{}

func (Substraction) Apply(lval, rval int) int {
	return lval - rval
}

func main() {
	mult := NewOperation(Addition{})
	fmt.Println(mult.GetName())
	res := mult.Operate(3, 5)
	fmt.Println(res)

	mult = NewOperation(Substraction{})
	fmt.Println(mult.GetName())
	res = mult.Operate(3, 5)
	fmt.Println(res)
}
