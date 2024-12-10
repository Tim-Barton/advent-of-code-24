package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	total := 0
	for i := range prints {
		if assessRules(prints[i], rules) {
			total += midNumber(prints[i])
		}
	}

	fmt.Println(total)
}
