package main

import (
	"testing"
)

// test data names correct for Part 2 expectations
// taken from the advent of code example
var safeDec []int = []int{7, 6, 4, 2, 1}
var unsafeInc []int = []int{1, 2, 7, 8, 9}
var unsafeDec []int = []int{9, 7, 6, 2, 1}
var safeInc []int = []int{1, 3, 2, 4, 5}
var safeDec2 []int = []int{8, 6, 4, 4, 1}
var safeInc2 []int = []int{1, 3, 6, 7, 9}

func TestDecreaingDampner(t *testing.T) {

	testResult := decreasingWithDampner(safeDec)
	if testResult == false {
		t.Fail()
	}

	testResult = decreasingWithDampner(unsafeDec)
	if testResult == true {
		t.Fail()
	}

	testResult = decreasingWithDampner(safeDec2)
	if testResult == false {
		t.Fail()
	}
}

func TestIncreasingDampner(t *testing.T) {
	testResult := increasingWithDampner(unsafeInc)
	if testResult == true {
		t.Fail()
	}

	testResult = increasingWithDampner(safeInc)
	if testResult == false {
		t.Fail()
	}

	testResult = increasingWithDampner(safeInc2)
	if testResult == false {
		t.Fail()
	}
}
