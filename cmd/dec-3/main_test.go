package main

import (
	"testing"
)

func TestSplitCommand(t *testing.T) {
	testData := "mul(123,456)"
	num1, num2 := splitcommand(testData)
	if num1 != 123 {
		t.Fail()
	}
	if num2 != 456 {
		t.Fail()
	}
}

func TestHandleline(t *testing.T) {
	// from the question details as provided correct example
	testData := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	expectedValue := 161
	returnValue := handleLine(testData)

	if returnValue != expectedValue {
		t.Fail()
	}
}
