package main

import "fmt"

// Определение поведения объекта в зависимости от его состояния

type Lightswitch struct {
	on bool
}

func (l *Lightswitch) TurnOn() {
	if l.on {
		fmt.Println("already on")
		return
	}
	l.on = !l.on
}
func (l *Lightswitch) TurnOff() {
	if !l.on {
		fmt.Println("already off")
		return
	}
	l.on = !l.on
}

func main() {
	l := Lightswitch{}
	l.TurnOn()
	l.TurnOn()
	l.TurnOff()
	l.TurnOff()
}
