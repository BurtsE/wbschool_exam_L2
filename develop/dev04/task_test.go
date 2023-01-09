package main

import "testing"

func TestInTest(t *testing.T) {
	testSlice := []string{"кот", "пес", "Лис34"}
	testInput := []string{"кот", "компот", "ПЕС"}
	output := []bool{true, false, false}
	t.Logf("testing on %v: ", testSlice)
	for i, w := range testInput {
		t.Logf("testing %s", w)
		if inSet(w, testSlice) != output[i] {
			t.Errorf("expected: %t. got: %t", output[i], !output[i])
		}
	}
}

func TestSortStrings(t *testing.T) {
	testInput := []string{"прием", "абв", "яйцо", "здоровье"}
	testOutput := []string{"еимпр", "абв", "йоця", "вдезоорь"}
	t.Logf("Testing on %s", testInput)
	for i := range testInput {
		t.Logf("testing %s", testInput[i])
		output := sortString(testInput[i])
		if output != testOutput[i] {

			t.Errorf("Expected: %s. Got: %s", testOutput[i], output)
		}
	}
}

func TestRefactorMap(t *testing.T) {
	testMap := map[string][]string{
		"a": []string{"абв", "бвг", "жвк"},
		"b": []string{"жвк", "бвг", "абв"},
		"c": []string{"кот", "бвг", "дом"},
		"d": []string{"нора", "рот", "жвк"},
	}
	testOutput := map[string][]string{
		"абв":  []string{"абв", "бвг", "жвк"},
		"жвк":  []string{"абв", "бвг", "жвк"},
		"кот":  []string{"бвг", "дом", "кот"},
		"нора": []string{"жвк", "нора", "рот"},
	}
	output := refactorMap(testMap)
	for key, val := range output {
		v, ok := testOutput[key]
		if !ok {
			t.Fatalf("Wrong key: %s", key)
		}
		for k := range val {
			if k >= len(v) || val[k] != v[k] {
				t.Errorf("Expected: %v. Got: %v", v, val)
			}
		}
	}
}

func TestCreateBook(t *testing.T) {
	testStrings := []string{"пятак", "пЯтка", "Тяпка", "листок", "слиток", "столик", "янтарь", "я"}
	testOutput := map[string][]string{
		"листок": []string{"листок", "слиток", "столик"},
		"пятак":  []string{"пятак", "пятка", "тяпка"},
		"янтарь": []string{"янтарь"},
	}
	output := createBook(testStrings)
	for key, val := range output {
		v, ok := testOutput[key]
		if !ok {
			t.Fatalf("Wrong key: %s", key)
		}
		for k := range val {
			if k >= len(v) || val[k] != v[k] {
				t.Errorf("Expected: %v. Got: %v", v, val)
			}
		}
	}
}
