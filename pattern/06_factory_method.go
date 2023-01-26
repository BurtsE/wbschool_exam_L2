package main

import "fmt"

// Паттерн Factory Method относится к порождающим паттернам
// Нужно определить интерфейс для создания объекта, но оставить подклассам решение вопроса о том, какой класс инстанцировать
// Используется, когда классу заранее неизвестно, объекты каких подклассов нужно создавать

// Интерфейс объектов, производимых фабрикой
type Calc interface {
	Add()
	Sub()
}

type Fabric struct {
}

func (f *Fabric) CreateCalc(t string) Calc {
	switch t {
	case "Simple":
		s := SimpleCalc{}
		return s.Init()
	case "Science":
		s := ScienceCalc{}
		return s.Init()
	default:
		s := SimpleCalc{}
		return s.Init()
	}
}

type ScienceCalc struct {
	params []float32
}

func (c *ScienceCalc) Init() *ScienceCalc {
	c.params = make([]float32, 4)
	c.params[0] = 3
	return c
}
func (c *ScienceCalc) Add() {}
func (c *ScienceCalc) Sub() {}

type SimpleCalc struct {
	params []int
}

func (c *SimpleCalc) Init() *SimpleCalc {
	c.params = make([]int, 8)
	c.params[0] = 5
	return c
}
func (c *SimpleCalc) Add() {}
func (c *SimpleCalc) Sub() {}

type AbstractFabric interface {
	CreateCalc(string) Calc
}

func main() {
	var f AbstractFabric
	f = new(Fabric)
	calc1 := f.CreateCalc("Simple")
	calc2 := f.CreateCalc("Science")
	fmt.Println(calc1, calc2)

}
