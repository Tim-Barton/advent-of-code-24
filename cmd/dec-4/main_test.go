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

func TestCheckInput2(t *testing.T) {
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
	expectedResult := 9
	result := checkInput2(testData)
	if expectedResult != result {
		fmt.Printf("Fail. Expected: %d, Returned: %d\n", expectedResult, result)
		t.Fail()
	}
}
