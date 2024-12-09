package main

import (
	"fmt"
	"testing"
)

func TestCheckInput(t *testing.T) {
	testData := []string{"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX"}
	expectedResult := 18
	result := checkInput(testData)
	if expectedResult != result {
		fmt.Printf("Fail. Expected: %d, Returned: %d\n", expectedResult, result)
		t.Fail()
	}
}
