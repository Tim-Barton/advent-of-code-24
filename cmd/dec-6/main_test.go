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

func TestObstructionFinder(t *testing.T) {
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

	expectedResult := 6

	lab, guard := parseInput(testData)
	result := FindObstructionPoints(guard, lab)

	if expectedResult != result {
		fmt.Printf("Expected: %d, Result: %d", expectedResult, result)
		t.Fail()
	}
}
