package main

import (
	"fmt"
	"math"
)

// Паттерн Visitor относится к поведенческим паттернам
// Шаблон позволяет расширить функционал класса или структуры без модификации их кода
// Применяется в случаях, когда фукнционал необходимо изменять часто, а объектов для изменения несколько

// В каждый объект встраивается метод visit(), вызывающий нужный метод у посетителя
type circle struct {
	x, y, r int
}

func (obj *circle) accept(v visitor) {
	v.visitForCircle(obj)
}

type square struct {
	x, y, w, h int
}

func (obj *square) accept(v visitor) {
	v.visitForSquare(obj)
}

// Общий интерфейс для всех посетителей
type visitor interface {
	visitForSquare(*square)
	visitForCircle(*circle)
}

// Каждый посетитель реализует свой функционал

type areaCalculator struct {
	area float64
}

func (a *areaCalculator) visitForSquare(s *square) {
	// Подсчет площади круга и присваивание значения переменной area
	a.area = float64(s.w * s.h)
}
func (a *areaCalculator) visitForCircle(s *circle) {
	// Подсчет площади круга и присваивание значения переменной area
	a.area = math.Pi * float64(s.r*s.r)
}

type perimeterCalculator struct {
	p int
}

func (a *perimeterCalculator) visitForSquare(s *square) {
	// Расчет периметра и присваивание значения переменной p
}

func (a *perimeterCalculator) visitForCircle(c *circle) {
	// Расчет периметра и присваивание значения переменной p
}

func main() {
	sq := &square{2, 2, 2, 2}
	c := &circle{2, 2, 2}
	ac := &areaCalculator{}
	sq.accept(ac)
	fmt.Println(ac.area)
	c.accept(ac)
	fmt.Println(ac.area)
}
