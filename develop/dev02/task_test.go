package main

import (
	"testing"
)

func TestEscape(t *testing.T) {
	testStrings := []string{"", "a4bc2d5e", "abcd", "45", `qwe\4\5`, `qwe\45`, `qwe\\5`, `\`, `sd3\`}
	testOutput := []string{"", "aaaabccddddde", "abcd", "", `qwe45`, `qwe44444`, `qwe\\\\\`, "", ""}
	for i := range testStrings {
		t.Logf("Testing: '%s'\n", testStrings[i])
		str, _ := expand(testStrings[i])
		if str != testOutput[i] {
			t.Errorf("Wrong output. Expected: '%s', got: '%s'\n", testOutput[i], str)
		}
	}
}

func TestEscapeError(t *testing.T) {
	testValidStrings := []string{"", "a4bc2d5e", "abcd", `qwe\4\5`, `qwe\45`, `qwe\\5`}
	testInvalidStrings := []string{"45", `\`, `sd3\`}

	for i := range testValidStrings {
		t.Logf("Testing: '%s'\n", testValidStrings[i])
		_, err := expand(testValidStrings[i])
		if err != nil {
			t.Errorf("Wrong output. Expected: <nil>, got: '%s'\n", err)
		}
	}
	for i := range testInvalidStrings {
		t.Logf("Testing: '%s'\n", testInvalidStrings[i])
		_, err := expand(testInvalidStrings[i])
		if err == nil {
			t.Errorf("Wrong output. Expected: error, got: '%s'\n", err)
		}
	}
}
