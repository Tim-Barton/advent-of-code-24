package main

import (
	"testing"
)

func TestRulesAssess(t *testing.T) {
	testRules := []Rule{
		Rule{"47", "53"},
		Rule{"97", "13"},
		Rule{"97", "61"},
		Rule{"97", "47"},
		Rule{"75", "29"},
		Rule{"61", "13"},
		Rule{"75", "53"},
		Rule{"29", "13"},
		Rule{"97", "29"},
		Rule{"53", "29"},
		Rule{"61", "53"},
		Rule{"97", "53"},
		Rule{"61", "29"},
		Rule{"47", "13"},
		Rule{"75", "47"},
		Rule{"97", "75"},
		Rule{"47", "61"},
		Rule{"75", "61"},
		Rule{"47", "29"},
		Rule{"75", "13"},
		Rule{"53", "13"},
	}

	testPrintOrder := [][]string{
		[]string{"75", "47", "61", "53", "29"},
		[]string{"97", "61", "53", "29", "13"},
		[]string{"75", "29", "13"},
		[]string{"75", "97", "47", "61", "53"},
		[]string{"61", "13", "29"},
		[]string{"97", "13", "75", "29", "47"},
	}

	response := assessRules(testPrintOrder[0], testRules)
	if !response {
		t.Fail()
	}
	response = assessRules(testPrintOrder[1], testRules)
	if !response {
		t.Fail()
	}
	response = assessRules(testPrintOrder[2], testRules)
	if !response {
		t.Fail()
	}
	response = assessRules(testPrintOrder[3], testRules)
	if response {
		t.Fail()
	}
	response = assessRules(testPrintOrder[4], testRules)
	if response {
		t.Fail()
	}
	response = assessRules(testPrintOrder[5], testRules)
	if response {
		t.Fail()
	}
}

func TestMidPoint(t *testing.T) {

	testData := []string{"75", "47", "61", "53", "29"}
	expected := 61
	mid := midNumber(testData)
	if expected != mid {
		t.Fail()
	}

}
