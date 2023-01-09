package main

import "fmt"

// Реализован в пакете http (Handler interface) - пример

type task struct {
	name string
}

type handler interface {
	Handle(*task)
	setNext(handler)
}

type Head struct {
	next handler
}

func (h *Head) Handle(t *task) {
	fmt.Println("Head handler...")
	if t.name != "Head" {
		h.next.Handle(t)
		return
	}
	fmt.Println(t.name)
}

func (h *Head) setNext(next handler) {
	h.next = next
}
type Body struct {
	next handler
}

func (h *Body) Handle(t *task) {
	fmt.Println("Body handler...")
	if t.name != "Body" {
		h.next.Handle(t)
		return
	}
	fmt.Println(t.name)
}

func (h *Body) setNext(next handler) {
	h.next = next
}
type Foot struct {
	next handler
}

func (h *Foot) Handle(t *task) {
	fmt.Println("Foot handler...")
	if t.name != "Foot" {
		h.next.Handle(t)
		return
	}
	fmt.Println(t.name)
}

func (h *Foot) setNext(next handler) {
	h.next = next
}

func main() {
	var tasks = []*task{
		&task{"Foot"},
		&task{"Body"},
		&task{"Foot"},
		&task{"Head"},
	}
	h := &Head{}
	b := &Body{}
	f := &Foot{}
	h.setNext(b)
	b.setNext(f)
	for _, t := range tasks {
		h.Handle(t)
	}
}
