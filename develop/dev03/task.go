package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки в файле по аналогии с консольной
утилитой sort (man sort — смотрим описание и основные
параметры): на входе подается файл из несортированными
строками, на выходе — файл с отсортированными.

Реализовать поддержку утилитой следующих ключей:

-k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Реализовать поддержку утилитой следующих ключей:

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учетом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	var k int
	var n, r, u, b bool
	var strs []string
	flag.IntVar(&k, "k", 0, "колонка для сортировки")
	flag.BoolVar(&n, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&r, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&u, "u", false, "не выводить повторяющиеся строки")
	flag.BoolVar(&b, "b", false, "игнорировать хвостовые пробелы")
	flag.Parse()

	strs = read(flag.Arg(0), u)

	ans := sortStrings(strs, k, n, r, b)
	for _, line := range ans {
		fmt.Println(line)
	}
}

func sortStrings(strs []string, k int, n, r, b bool) []string {
	var x, y int
	newStrings := makeTable(strs, k, b)
	x, y = k-1, k-1
	if k == 0 {
		x, y = 0, 0
	}
	sort.SliceStable(newStrings, func(i, j int) bool {
		if k >= len(newStrings[i]) {
			x = len(newStrings[i]) - 1
		}
		if k >= len(newStrings[j]) {
			y = len(newStrings[j]) - 1
		}
		if n {
			p, _ := strconv.ParseFloat(strings.TrimSpace(strs[x]), 64)
			q, _ := strconv.ParseFloat(strings.TrimSpace(strs[y]), 64)
			return p < q
		}
		return newStrings[i][x] < newStrings[j][y]
	})

	for i := range strs {
		strs[i] = strings.Join(newStrings[i], " ")
	}
	if r {
		for i := 0; i < len(strs)/2; i++ {
			strs[i], strs[len(strs)-1] = strs[len(strs)-1], strs[i]
		}
	}
	return strs
}

func abs(num int) int {
	if num < 0 {
		num = -num
	}
	return num
}
func makeTable(strs []string, k int, b bool) (table [][]string) {
	table = make([][]string, len(strs))
	if k <= 0 {
		for i := range strs {
			table[i] = []string{strs[i]}
		}
		return
	}
	for i := range strs {
		if b {
			table[i] = strings.Fields(strs[i])
		} else {
			table[i] = strings.Split(strs[i], " ")
		}
	}
	return
}

func read(filename string, u bool) []string {
	var err error
	var strs []string
	var m = make(map[string]int)
	file := os.Stdin
	if filename != "" {
		file, err = os.Open(filename)
		if err != nil {
			log.Fatal("cannot open file", err)
		}
		defer file.Close()
	}
	s := bufio.NewScanner(file)
	for s.Scan() {
		m[s.Text()]++
	}
	for key, val := range m {
		if u {
			strs = append(strs, key)
		} else {
			for i := 0; i < val; i++ {
				strs = append(strs, key)
			}
		}
	}
	return strs
}
