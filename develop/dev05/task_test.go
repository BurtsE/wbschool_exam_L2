package main

import (
	"bufio"
	"os"
	"testing"
)

func TestBeforeN(t *testing.T) {
	testInput := [][]int{
		[]int{5, 8},
		[]int{5, 5},
		[]int{5, 2},
		[]int{3, 1},
	}
	testOutput := []int{0, 0, 3, 2}
	for i, input := range testInput {
		output := BeforeN(input[0], input[1])
		if output != testOutput[i] {
			t.Errorf("Expected: %d.Got: %d", testOutput[i], output)
		}
	}
}
func TestAfterN(t *testing.T) {
	var A = 6
	testInput := [][]string{
		make([]string, 7),
		make([]string, 0),
		make([]string, 3),
		make([]string, 4),
	}
	testOutput := []int{6, 0, 2, 3}
	for i, input := range testInput {
		output := AfterN(A, input)
		if output != testOutput[i] {
			t.Errorf("Expected: %d.Got: %d", testOutput[i], output)
		}
	}
}

func TestMatchesNums(t *testing.T) {
	var i, v, f = false, false, false
	pattern := `^\d*$`
	testStrings := []string{"432", "5436", "9879", "123", "54", ""}
	for _, str := range testStrings {
		t.Logf("testing %s", str)
		if !matches(pattern, str, i, v, f) {
			t.Fatal("string should match")
		}
		v = !v
		if matches(pattern, str, i, v, f) {
			t.Fatal("string shouldn't match. flag v is ignored")
		}
		v = !v
		f = !f
		if matches(pattern, str, i, v, f) {
			t.Fatal("string shouldn't match. flag f might being ignored")
		}
		f = !f
	}
}
func TestMatchesStrings(t *testing.T) {
	var i, v, f = true, false, false
	pattern := `^[a-z]*$`
	testStrings := []string{"Xv", "fdAv", "ASD", "ergGDe", "czVd"}
	for _, str := range testStrings {
		t.Logf("testing %s", str)
		if !matches(pattern, str, i, v, f) {
			t.Fatal("string should match")
		}
		v = !v
		if matches(pattern, str, i, v, f) {
			t.Fatal("string shouldn't match. flag v is ignored")
		}
		v = !v
		f = !f
		if matches(pattern, str, i, v, f) {
			t.Fatal("string shouldn't match. flag f might being ignored")
		}
		f = !f
		i = !i
		if matches(pattern, str, i, v, f) {
			t.Fatal("string shouldn't match. flag i might being ignored")
		}
		i = !i
	}
}

func TestGrep1(t *testing.T) {
	file, _ := os.Create("/tmp/yourfile")
	w := bufio.NewWriter(file)
	testInput := []string{"кот", "файл", "решето", "123435", "файл", "йл", "от", "123"}
	patterns := []string{"от"}
	testOutput := []string{"кот", "йл", "от"}
	A, B := 0, 1
	c, f, v, i, n := false, false, false, false, false
	grep(A, B, c, f, v, i, n, patterns, testInput, w)
	w.Flush()
	file.Close()
	file, _ = os.Open("/tmp/yourfile")
	defer file.Close()
	r := bufio.NewScanner(file)
	var q int
	for r.Scan() {
		if r.Text() != testOutput[q] {
			t.Errorf("wrong: %s", r.Text())
		}
		q++
	}
}

func TestGrep2(t *testing.T) {
	file, _ := os.Create("/tmp/yourfile")

	w := bufio.NewWriter(file)
	testInput := []string{"кот", "файл", "решето", "123435", "файл2", "йл", "от", "123"}
	patterns := []string{"фай"}
	testOutput := []string{"файл", "файл2"}
	A, B := 0, 0
	c, f, v, i, n := false, false, false, false, false
	grep(A, B, c, i, v, f, n, patterns, testInput, w)
	w.Flush()
	file.Close()
	file, _ = os.Open("/tmp/yourfile")
	defer file.Close()
	r := bufio.NewScanner(file)
	var q int
	for r.Scan() {
		if q >= len(testOutput) || r.Text() != testOutput[q] {
			t.Errorf("wrong: %s", r.Text())
		}
		q++
	}
}

func TestGrep3(t *testing.T) {
	file, _ := os.Create("/tmp/yourfile")
	w := bufio.NewWriter(file)
	testInput := []string{"кот", "файл", "решето", "123435", "файл2", "йл", "от", "123"}
	patterns := []string{"от"}
	testOutput := []string{"йл", "от", "123"}
	A, B := 1, 1
	c, f, v, i, n := false, true, false, false, false
	grep(A, B, c, i, v, f, n, patterns, testInput, w)
	w.Flush()
	file.Close()
	file, _ = os.Open("/tmp/yourfile")
	defer file.Close()
	r := bufio.NewScanner(file)
	var q int
	for r.Scan() {
		if q >= len(testOutput) || r.Text() != testOutput[q] {
			t.Errorf("wrong: %s", r.Text())
		}
		q++
	}
}
