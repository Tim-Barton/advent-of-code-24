package main

import (
	"fmt"
	"testing"
)

func TestNewCalibration(t *testing.T) {
	testData := "190: 10 19"
	cal := NewCalibration(testData)

	if cal.answer != 190 {
		t.Fail()
	}

	if len(cal.inputs) != 2 {
		t.Fail()
	}

}

func TestAssessCalibration(t *testing.T) {
	testData := "190: 10 19\n" +
		"3267: 81 40 27\n" +
		"83: 17 5\n" +
		"156: 15 6\n" +
		"7290: 6 8 6 15\n" +
		"161011: 16 10 13\n" +
		"192: 17 8 14\n" +
		"21037: 9 7 18 13\n" +
		"292: 11 6 16 20\n"

	expectedResult := 3749
	result := AssessCalibrationData(testData)

	if expectedResult != result {
		fmt.Printf("Expected: %d, Returned: %d\n", expectedResult, result)
		t.Fail()
	}
}

func TestAssessCalibration2(t *testing.T) {
	testData := "190: 10 19\n" +
		"3267: 81 40 27\n" +
		"83: 17 5\n" +
		"156: 15 6\n" +
		"7290: 6 8 6 15\n" +
		"161011: 16 10 13\n" +
		"192: 17 8 14\n" +
		"21037: 9 7 18 13\n" +
		"292: 11 6 16 20\n"

	expectedResult := 11387
	result := AssessCalibrationData2(testData)

	if expectedResult != result {
		fmt.Printf("Expected: %d, Returned: %d\n", expectedResult, result)
		t.Fail()
	}
}
