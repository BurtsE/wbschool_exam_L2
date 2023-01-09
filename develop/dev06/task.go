package main

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

	file := os.Stdin
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		columns := strings.Split(scanner.Text(), d)
		if len(columns) > 1 {
			for _, n := range fields {
				if n < len(columns)+1 {
					fmt.Print(columns[n-1], " ")
				}
			}
			fmt.Print("\n")
		} else if !s {
			fmt.Println(columns[0])
		}
	}
}
