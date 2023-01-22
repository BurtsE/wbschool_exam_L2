package main

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Объявление и парсинг флагов
	var s bool
	var f, d string
	flag.StringVar(&f, "f", "", "выбрать поля")
	flag.StringVar(&d, "d", ".", "использовать другой разделитель")
	flag.BoolVar(&s, "s", false, "только строки с разделителем")
	flag.Parse()
	fieldsStr := strings.Split(f, ",")
	fields := make([]int, len(fieldsStr)) // Номера столбцов для вывода
	var err error
	for i := range fields {
		fields[i], err = strconv.Atoi(fieldsStr[i])
		if err != nil {
			log.Fatalf("must specify fields with -f: %s", f)
		}
	}

	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	Cut(s, fields, d, scanner, writer)
}
func Cut(s bool, fields []int, d string, scanner *bufio.Scanner, w *bufio.Writer) {

	for scanner.Scan() {
		columns := strings.Split(scanner.Text(), d)
		if len(columns) > 1 {
			for i, n := range fields {
				if n < len(columns)+1 {
					if i == 0 {
						fmt.Fprintf(w, "%s", columns[n-1])
					} else {
						fmt.Fprintf(w, ":%s", columns[n-1])
					}

				}
			}
			fmt.Fprint(w, "\n")
		} else if !s {
			fmt.Fprintf(w, "%s\n", columns[0])
		}
	}
}
