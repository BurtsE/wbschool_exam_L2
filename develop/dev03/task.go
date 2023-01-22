package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	var k int
	var n, r, u bool
	var strs []string
	flag.IntVar(&k, "k", -1, "колонка для сортировки")
	flag.BoolVar(&n, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&r, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&u, "u", false, "не выводить повторяющиеся строки")
	flag.Parse()
	strs = read(flag.Arg(0), u)

	ans := sortStrings(strs, -1, r)
	for _, line := range ans {
		fmt.Println(line)
	}
}

func sortStrings(strs []string, k int, r bool) []string {
	if k == -1 {
		sort.SliceStable(strs, func(i, j int) bool {
			return strs[i] < strs[j]
		})
		return strs
	}
	newStrings := make([][]string, len(strs))
	for i := range strs {
		newStrings[i] = strings.Split(strs[i], " ")
	}
	var a = k
	var b = k
	sort.SliceStable(newStrings, func(i, j int) bool {
		if a >= len(newStrings[i]) {
			a = len(newStrings[i]) - 1
		}
		if b >= len(newStrings[j]) {
			b = len(newStrings[j]) - 1
		}
		return newStrings[a][i] < newStrings[b][j]
	})
	for i := range strs {
		strs[i] = strings.Join(newStrings[i], " ")
	}
	if r {
		for i :=0; i < len(strs)/2; i++ {
			strs[i], strs[len(strs)-1] = strs[len(strs)-1], strs[i]
		}
	}
	return strs
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
