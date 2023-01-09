package main

import (
	"fmt"
)

func main() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4]
	fmt.Println(b)
}

// Программа напечатет слайс , состоящий из элементов с индексами 1, 2, 3 слайса a: {77, 78, 79}
