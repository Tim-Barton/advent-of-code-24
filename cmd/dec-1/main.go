package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func splitLine(s string) (string, string) {
	parts := strings.Fields(s)
	return parts[0], parts[1]
}

func part1(leftlines, rightlines []int) int {
	total := 0
	for i := range leftlines {
		left := leftlines[i]
		right := rightlines[i]
		if left > right {
			total += (left - right)
		} else {
			total += (right - left)
		}
	}

	return total
}

func part2(leftlines, rightlines []int) int {
	total := 0
	for i := range leftlines {
		left := leftlines[i]
		count := 0
		for j := range rightlines {
			if left == rightlines[j] {
				count += 1
			}
		}
		total += (left * count)
	}

	return total
}

func main() {
	input, _ := os.Open("input")
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	leftlines := []int{}
	rightlines := []int{}

	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		left, right := splitLine(scanner.Text())
		leftInt, err := strconv.Atoi(left)
		if err != nil {
			fmt.Printf("Left is bad on line %s\n", scanner.Text())
			continue
		}
		rightInt, err := strconv.Atoi(right)
		if err != nil {
			fmt.Printf("Right is bad on line %s\n", scanner.Text())
			continue
		}
		leftlines = append(leftlines, leftInt)
		rightlines = append(rightlines, rightInt)
	}

	slices.Sort(leftlines)
	slices.Sort(rightlines)

	if len(leftlines) != len(rightlines) {
		fmt.Println("List are differnet lengths")
		os.Exit(1)
	}

	part1Total := part1(leftlines, rightlines)
	part2Total := part2(leftlines, rightlines)

	fmt.Printf("Part 1: %d\n", part1Total)
	fmt.Printf("Part 2: %d\n", part2Total)

}
