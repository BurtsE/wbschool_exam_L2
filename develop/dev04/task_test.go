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
