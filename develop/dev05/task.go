package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
)

func Usage() {
	fmt.Printf("Usage: main [flags] PATTERNS [FILE...]\n")
}

func main() {
	// Объявление и парсинг флагов
	var c, i, v, f, n bool
	var A, B, C int
	flag.IntVar(&A, "A", 0, "печатать +N строк после совпадения")
	flag.IntVar(&B, "B", 0, "печатать +N строк до совпадения")
	flag.IntVar(&C, "C", 0, "печатать ±N строк вокруг совпадения")
	flag.BoolVar(&c, "c", false, "печатать количество строк")
	flag.BoolVar(&i, "i", false, "игнорировать регистр")
	flag.BoolVar(&v, "v", false, "вместо совпадения, исключать ")
	flag.BoolVar(&f, "f", false, "точное совпадение со строкой, не паттерн")
	flag.BoolVar(&n, "n", false, "напечатать номер строки")
	flag.Parse()
	if C > A {
		A = C
	}
	if C > B {
		B = C
	}

	// Проверка кол-ва аргументов
	length := len(flag.Args())
	if length < 1 {
		Usage()
		log.Fatalf("Not enough arguments: %v", flag.Args())
	}

	// Попытка открыть файл. Если не удалась, последний аргумент считается паттерном
	patterns := flag.Args()[:length-1]
	file, err := os.Open(flag.Arg(length - 1))
	if err != nil {
		patterns = append(patterns, flag.Arg(length-1))
		file = os.Stdin
		log.Println(err)
	}

	// Считываем файл в слайс
	scanner := bufio.NewScanner(file)
	var lines = make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var set = make(map[string]bool) // Множество для предотвращения многократной печати одной и той же строки.
	var answer = make([]string, 0)  // Результат работы
	var answerLine = make([]int, 0) // Номера строк в результате

	for num := 0; num < len(lines); num++ { // Проверка каждой строки
		for m := 0; m < len(patterns); m++ { // Соответствие каждому из паттернов
			if matches(patterns[m], lines[num], i, v, f) {
				start := BeforeN(num, B)
				end := num + AfterN(A, lines[num:])
				for q := start; q < end; q++ {
					if _, ok := set[lines[q]]; !ok { // Если строка еще не добавлена в ответ
						set[lines[q]] = false
						answer = append(answer, lines[q])
						answerLine = append(answerLine, q)
					}
				}
				set[lines[num]] = true // True у строк, подходящих паттерну
				//num = end // Пропуск строк, добавленных в ответ
			}
		}
	}
	// Размер результата
	if c {
		fmt.Println(len(answer))
		return
	}
	for i, line := range answer {
		if n {
			if set[line] {
				fmt.Printf("%d: ", answerLine[i])
			} else {
				fmt.Printf("%d- ", answerLine[i])
			}
		}
		fmt.Println(line)
	}
	//fmt.Println(A, B, C, c, i, v, f, n)
}

// Вычисляет смещение на -B с учетом границы массива
func BeforeN(num, B int) int {
	if num > B {
		return num - B
	} else {
		return 0
	}
}

// Вычисляет смещение на +A с учетом границы массива
func AfterN(A int, afterLines []string) int {
	var end int
	for i := 0; i < A+2 && i < len(afterLines); i++ {
		end = i
	}
	return end
}

// Фукнция проверяет, подходит ли строка для ответа с учетом флагов
func matches(pattern, s string, i, v, f bool) (answer bool) {
	defer func(invert bool) {
		if v {
			answer = !answer
		}
	}(v)
	if f {
		answer = (pattern == s)
		return
	}
	if i {
		pattern = "(?i)" + pattern
	}
	exp, _ := regexp.Compile(pattern)

	if exp.FindStringSubmatch(s) != nil {
		answer = true
	}
	return
}
