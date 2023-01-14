package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

func Usage() {
	log.Printf("Usage: main [flags] PATTERNS [FILE...]\n")
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
		os.Exit(1)
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

	grep(A, B, c, i, v, f, n, patterns, lines, os.Stdout)
}

func grep(A, B int, c, i, v, f, n bool, patterns []string, lines []string, w io.Writer) {

	var set = make(map[int]bool)    // Множество для предотвращения многократной печати одной и той же строки.
	var answer = make([]string, 0)  // Результат работы
	var answerLine = make([]int, 0) // Номера строк в результате 123

	for num := 0; num < len(lines); num++ { // Проверка каждой строки
		for m := 0; m < len(patterns); m++ { // Соответствие каждому из паттернов
			if matches(patterns[m], lines[num], i, v, f) {
				start := BeforeN(num, B)
				end := num + AfterN(A, lines[num:])
				for q := start; q <= end; q++ {
					if _, ok := set[q]; !ok { // Если строка еще не добавлена в ответ
						set[q] = false
						answer = append(answer, lines[q])
						answerLine = append(answerLine, q)
					}
				}
				set[num] = true // True у строк, подходящих паттерну
				//num = end // Пропуск строк, добавленных в ответ
			}
		}
	}
	// Размер результата
	if c {
		fmt.Fprintf(w, "%d \n", len(answer))
		return
	}
	for i, line := range answer {
		if n {
			if set[answerLine[i]] {
				fmt.Fprintf(w, "%d: ", answerLine[i])
			} else {
				fmt.Fprintf(w, "%d- ", answerLine[i])
			}
		}
		fmt.Fprintf(w, "%s\n", line)
	}
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
	if len(afterLines) == 0 {
		return 0
	}
	if A >= len(afterLines) {
		return len(afterLines) - 1
	}
	return A
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
