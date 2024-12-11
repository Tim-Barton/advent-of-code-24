package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Rule struct {
	first  string
	second string
}

func assessRules(printOrder []string, rules []Rule) bool {
	for i := range rules {
		if !assessRule(printOrder, rules[i]) {
			return false
		}
	}
	return true
}

func assessRule(printOrder []string, rule Rule) bool {
	firstIndex := -1
	secondIndex := -1
	for i := range printOrder {
		if printOrder[i] == rule.first {
			firstIndex = i
		}
		if printOrder[i] == rule.second {
			secondIndex = i
		}
	}
	if firstIndex != -1 && secondIndex != -1 && secondIndex < firstIndex {
		return false
	}
	return true
}

func fixOrder(printOrder []string, rules []Rule) []string {
	for !assessRules(printOrder, rules) {
		for i := range rules {
			correct := assessRule(printOrder, rules[i])
			if !correct {
				firstIndex := slices.Index(printOrder, rules[i].first)
				secondIndex := slices.Index(printOrder, rules[i].second)
				printOrder = slices.Delete(printOrder, firstIndex, firstIndex+1)
				printOrder = slices.Insert(printOrder, secondIndex, rules[i].first)
			}
		}
	}

	return printOrder
}

func midNumber(printOrder []string) int {

	midPoint := (len(printOrder) / 2)
	midValue, _ := strconv.Atoi(printOrder[midPoint])
	return midValue
}

func main() {

	input, _ := os.Open("input")
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	rules := []Rule{}
	prints := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if match, _ := regexp.MatchString(`\d+\|\d+`, line); match {
			parts := strings.Split(line, "|")
			rules = append(rules, Rule{first: parts[0], second: parts[1]})
		} else if match, _ := regexp.MatchString(`\d+\,\d+`, line); match {
			prints = append(prints, strings.Split(line, ","))
		}
	}

	goodTotal := 0
	correctedTotal := 0
	for i := range prints {
		if assessRules(prints[i], rules) {
			goodTotal += midNumber(prints[i])
		} else {
			newOrder := fixOrder(prints[i], rules)
			correctedTotal += midNumber(newOrder)
		}
	}

	fmt.Println(goodTotal)
	fmt.Println(correctedTotal)
}
