package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const maxChange int = 3

func increasing(input []int) bool {
	for i := 0; i < len(input)-1; i++ {
		if input[i] > input[i+1] {
			return false
		}

		if input[i] == input[i+1] {
			return false
		}

		if (input[i+1] - input[i]) > maxChange {
			return false
		}
	}
	return true
}

func increasingWithDampner(input []int) bool {
	for i := 0; i < len(input)-1; i++ {
		if input[i] > input[i+1] {
			return increasing(append(input[:i+1], input[i+2:]...))
		}

		if input[i] == input[i+1] {
			return increasing(append(input[:i+1], input[i+2:]...))
		}

		if (input[i+1] - input[i]) > maxChange {
			return increasing(append(input[:i+1], input[i+2:]...))
		}
	}
	return true
}

func decreasing(input []int) bool {
	fmt.Printf("Determining Dec on array %v\n", input)
	for i := 0; i < len(input)-1; i++ {
		if input[i] < input[i+1] {
			return false
		}

		if input[i] == input[i+1] {
			return false
		}

		if (input[i] - input[i+1]) > maxChange {
			return false
		}
	}
	return true
}

func decreasingWithDampner(input []int) bool {
	for i := 0; i < len(input)-1; i++ {
		if input[i] < input[i+1] {
			fmt.Printf("Using Dec Dampner on array %v\n", input)
			return decreasing(append(input[:i+1], input[i+2:]...))
		}

		if input[i] == input[i+1] {
			fmt.Printf("Using Dec Dampner on array %v\n", input)
			return decreasing(append(input[:i+1], input[i+2:]...))
		}

		if (input[i] - input[i+1]) > maxChange {
			fmt.Printf("Using Dec Dampner on array %v\n", input)
			return decreasing(append(input[:i+1], input[i+2:]...))
		}
	}
	return true
}

func splitLine(inputLine string) []int {
	parts := strings.Fields(inputLine)
	retVal := []int{}
	for i := range parts {
		intVal, err := strconv.Atoi(parts[i])
		if err != nil {
			fmt.Printf("Value is bad on line %s\n", inputLine)
			continue
		}
		retVal = append(retVal, intVal)
	}
	return retVal
}

func main() {

	input, _ := os.Open("input")
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	safeReports := 0
	safeReports2 := 0

	for scanner.Scan() {
		line := splitLine(scanner.Text())
		isSafe := false
		if line[0] > line[1] {
			isSafe = decreasing(line)
		} else {
			isSafe = increasing(line)
		}
		if isSafe {
			safeReports += 1
		}

		isSafe2 := false
		if line[0] > line[1] {
			isSafe2 = decreasingWithDampner(line)
			fmt.Printf("Result: %v\n", isSafe)
			fmt.Printf("Result2: %v\n", isSafe2)
		} else {
			isSafe2 = increasingWithDampner(line)
		}

		if isSafe2 {
			safeReports2 += 1
		}
	}

	fmt.Printf("Part 1 asnwer: %d\n", safeReports)
	fmt.Printf("Part 2 answer: %d\n", safeReports2)

}
