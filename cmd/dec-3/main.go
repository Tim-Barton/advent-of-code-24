package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func splitcommand(input string) (int, int) {
	substr, _ := strings.CutPrefix(input, "mul(")
	substr, _ = strings.CutSuffix(substr, ")")
	values := strings.Split(substr, ",")
	num1, _ := strconv.Atoi(values[0])
	num2, _ := strconv.Atoi(values[1])
	return num1, num2
}

func handleLine(input string) int {
	pattern := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	commands := pattern.FindAllString(input, -1)
	commandTotal := 0
	for i := range commands {
		num1, num2 := splitcommand(commands[i])
		sub := num1 * num2
		commandTotal += sub
	}
	return commandTotal
}

func handleLine2(input string, mulEnabled bool) (int, bool) {
	fmt.Println("New Line")
	pattern := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	dontString := "don't()"
	doString := "do()"
	maxMulString := "mul(123,456)"
	commandTotal := 0
	for inputIndex := 0; inputIndex < len(input); inputIndex++ {
		if mulEnabled {
			switch input[inputIndex] {
			case 'm':
				endIndex := inputIndex + len(maxMulString)
				if endIndex > len(input) {
					endIndex = len(input)
				}

				substr := input[inputIndex:endIndex]
				match := pattern.FindString(substr)
				if match != "" {
					fmt.Printf("Match: %s\n", match)
					num1, num2 := splitcommand(match)
					commandTotal += (num1 * num2)
					inputIndex += len(match) - 1
				}
				continue
			case 'd':
				substr := input[inputIndex : inputIndex+len(dontString)]
				if substr == dontString {
					fmt.Println("Disabled")
					mulEnabled = false
					inputIndex += len(dontString) - 1
					continue
				}

			}
		} else {
			switch input[inputIndex] {
			case 'd':
				substr := input[inputIndex : inputIndex+len(doString)]
				if substr == doString {
					fmt.Println("Enabled")
					mulEnabled = true
					inputIndex += len(doString) - 1
					continue
				}
			}
		}
	}
	return commandTotal, mulEnabled
}

func main() {
	input, _ := os.Open("input")
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	documentTotal, lineTotal := 0, 0
	enabled := true
	for scanner.Scan() {
		line := scanner.Text()
		lineTotal, enabled = handleLine2(line, enabled)
		documentTotal += lineTotal
	}

	fmt.Printf("Ouput: %d\n", documentTotal)
}
