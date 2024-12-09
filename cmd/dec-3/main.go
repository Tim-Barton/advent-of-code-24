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

func main() {
	input, _ := os.Open("input")
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	documentTotal := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineTotal := handleLine(line)
		documentTotal += lineTotal
	}

	fmt.Printf("Ouput: %d\n", documentTotal)
}
