package main

import (
	"fmt"
	"testing"
)

func TestImportExportLab(t *testing.T) {

	testData := "....#.....\n" +
		".........#\n" +
		"..........\n" +
		"..#.......\n" +
		".......#..\n" +
		"..........\n" +
		".#..^.....\n" +
		"........#.\n" +
		"#.........\n" +
		"......#...\n"

	lab, guard := parseInput(testData)
	//t.FailNow()
	guard.Patrol(lab)
	expectedResult := 41
	if expectedResult != guard.Steps() {
		fmt.Printf("Expected: %d, Result: %d", expectedResult, guard.Steps())
		t.Fail()
	}
}
