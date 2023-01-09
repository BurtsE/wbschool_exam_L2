package main

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	fmt.Println(expand(`\3`))
}

func expand(input string) (string, error) {
	str := []rune(input)
	var answer = make([]rune, 0)
	var start int
	var number int
	for i := 0; i < len(str); {
		if i < len(str) && unicode.IsLetter(str[i]) {
			i++
			start = i
		} else if i < len(str) && str[i] == '\\' {
			if i == len(str)-1 {
				return "", errors.New("wrong format")
			}
			i += 2
			start = i
		}
		for i < len(str) && unicode.IsDigit(str[i]) {
			if i == 0 {
				return "", errors.New("string starts with wrong symbol")
			}
			i++
		}

		number, err := strconv.Atoi(string(str[start:i]))
		if err != nil {
			number = 1
		}
		for k := 0; k < number; k++ {
			answer = append(answer, str[start-1])
		}
	}
	for k := 0; k < number; k++ {
		answer = append(answer, str[start-1])
	}
	return string(answer), nil
}
