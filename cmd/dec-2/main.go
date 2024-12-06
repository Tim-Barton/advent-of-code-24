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
			return decreasing(append(input[:i+1], input[i+2:]...))
		}

		if input[i] == input[i+1] {
			return decreasing(append(input[:i+1], input[i+2:]...))
		}

		if (input[i] - input[i+1]) > maxChange {
			return decreasing(append(input[:i+1], input[i+2:]...))
		}
	}
	return true
}

func assessReport(input []int, allowedErrors int) bool {
	isSafe := false
	if input[0] > input[1] {
		isSafe = decreasing(input)
	} else {
		isSafe = increasing(input)
	}

	if allowedErrors > 0 && isSafe == false {
		for i := range input {
			tmp := []int{}
			tmp = append(tmp, input[:i]...)
			tmp = append(tmp, input[i+1:]...)
			// if any variant of the report is safe (true) then return
			if assessReport(tmp, allowedErrors-1) {
				return true
			}
		}
		return isSafe
	} else {
		return isSafe
	}

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
		isSafe := assessReport(line, 0)
		if isSafe {
			safeReports += 1
		}

		isSafe2 := assessReport(line, 1)
		if isSafe2 {
			safeReports2 += 1
		}
	}

	fmt.Printf("Part 1 asnwer: %d\n", safeReports)
	fmt.Printf("Part 2 answer: %d\n", safeReports2)

}
