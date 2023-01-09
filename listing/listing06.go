package main

import (
	"fmt"
)

func main() {
	var s = []string{"1", "2", "3"}
	modifySlice(s)
	fmt.Println(s)
}
func modifySlice(i []string) {
	i[0] = "3"
	i = append(i, "4")
	i[1] = "5"
	i = append(i, "6")
}

/*
Вывод:
{3, 2, 3}

Если определить слайс таким образом, как в программе, его длинна len и вместимость cap будут равны числу элементов, переданных при инициализации.
Слайс является ссылочным типом данных, поэтому функция modifySlice, получив его копию, может изменять содержимое. Первый элемент слайса меняется на 3.
Сам слайс является ссылкой на массив. Функция append увеличивает длинну слайса и меняет следующий элемент массива на указанный ,если число
элементов слайса не привысит его cap (длинну массива). В противном случае создается новый массив с удвоенной длинной, в который копируются элементы
переданного слайса и значения, которые в этот слайс пытались вложить. Затем возвращается слайс с соответствующей длинной и удвоенным cap.
Т.к. изначально длинна и вместимость слайса равны, append присваивает переменной i указатель на новый массив, и последующие изменения применяются к нему.
Ссылка на исходный массив оказывается утеряна, и поскольку слайс в main указывает на него, этот слайс остается неизменным.
*/