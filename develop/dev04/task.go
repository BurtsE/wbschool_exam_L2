package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	strings := []string{"пятак", "пЯтка", "Тяпка", "листок", "слиток", "столик", "янтарь", "я", "ПиЛа", "лИпа", "Пост", "стоП"}
	fmt.Println(createBook(strings))
}

func createBook(input []string) map[string][]string {
	help := make(map[string][]string)
	for _, s := range input {
		if len([]rune(s)) > 1 {
			s = strings.ToLower(s)
			key := sortString(s)
			if _, ok := help[key]; !ok {
				help[key] = make([]string, 0)
			}
			if !inSet(s, help[key]) {
				help[key] = append(help[key], s)
			}
		}
	}
	ans := refactorMap(help)
	return ans
}

func refactorMap(m map[string][]string) map[string][]string {
	ans := make(map[string][]string)
	for _, val := range m {
		ans[val[0]] = val
		sort.Strings(ans[val[0]])
	}
	return ans
}

func sortString(str string) string {
	arr := []rune(str)
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	return string(arr)
}

func inSet(s string, set []string) bool {
	for _, str := range set {
		if s == str {
			return true
		}
	}
	return false
}
