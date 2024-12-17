package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Operator int

const (
	Add Operator = iota
	Multiply
	Concat
)

var stateName = map[Operator]string{
	Add:      "+",
	Multiply: "*",
	Concat:   "||",
}

func (pt Operator) String() string {
	return stateName[pt]
}

type Calibration struct {
	answer int
	inputs []int
}

func buildOperatorList(depth int) <-chan []Operator {
	output := make(chan []Operator)

	go func() {
		var recursion func([]Operator, int)

		recursion = func(current []Operator, depth int) {
			if depth == 1 {
				newList := append([]Operator{}, current...)
				output <- append(newList, Add)
				newList = append([]Operator{}, current...)
				output <- append(newList, Multiply)
			} else {
				newList := append([]Operator{}, current...)
				recursion(append(newList, Add), depth-1)
				newList = append([]Operator{}, current...)
				recursion(append(newList, Multiply), depth-1)
			}
		}
		recursion([]Operator{}, depth)
		close(output)
	}()

	return output
}

func buildOperatorList2(depth int) <-chan []Operator {
	output := make(chan []Operator)

	go func() {
		var recursion func([]Operator, int)

		recursion = func(current []Operator, depth int) {
			if depth == 1 {
				newList := append([]Operator{}, current...)
				output <- append(newList, Add)
				newList = append([]Operator{}, current...)
				output <- append(newList, Multiply)
				newList = append([]Operator{}, current...)
				output <- append(newList, Concat)
			} else {
				newList := append([]Operator{}, current...)
				recursion(append(newList, Add), depth-1)
				newList = append([]Operator{}, current...)
				recursion(append(newList, Multiply), depth-1)
				newList = append([]Operator{}, current...)
				recursion(append(newList, Concat), depth-1)
			}
		}
		recursion([]Operator{}, depth)
		close(output)
	}()

	return output
}

func (c Calibration) Assess() bool {
	assessments := buildOperatorList(len(c.inputs) - 1)

	for opList := range assessments {
		localAnswer := c.inputs[0]
		for i := 1; i < len(c.inputs); i++ {
			switch opList[i-1] {
			case Add:
				localAnswer += c.inputs[i]
			case Multiply:
				localAnswer *= c.inputs[i]
			}
		}
		if localAnswer == c.answer {
			//fmt.Printf("Answer: %d, inputs: %v, Operators: %v\n", c.answer, c.inputs, opList)
			return true
		}
	}
	return false
}

func (c Calibration) Assess2() bool {
	assessments := buildOperatorList2(len(c.inputs) - 1)

	for opList := range assessments {
		localAnswer := c.inputs[0]
		for i := 1; i < len(c.inputs); i++ {
			switch opList[i-1] {
			case Add:
				localAnswer += c.inputs[i]
			case Multiply:
				localAnswer *= c.inputs[i]
			case Concat:
				localAnswer, _ = strconv.Atoi(strconv.Itoa(localAnswer) + strconv.Itoa(c.inputs[i]))
			}
		}
		if localAnswer == c.answer {
			//fmt.Printf("Answer: %d, inputs: %v, Operators: %v\n", c.answer, c.inputs, opList)
			return true
		}
	}
	return false
}

func NewCalibration(s string) Calibration {
	base := strings.Split(s, ":")
	inputsStrings := strings.Split(strings.Trim(base[1], " "), " ")
	answer, _ := strconv.Atoi(base[0])
	inputs := []int{}
	for i := range inputsStrings {
		input, _ := strconv.Atoi(inputsStrings[i])
		inputs = append(inputs, input)
	}
	return Calibration{answer: answer, inputs: inputs}
}

func parseInput(input string) []Calibration {
	lines := strings.Split(input, "\n")
	calibrations := []Calibration{}
	for i := range lines {
		if lines[i] != "" {
			calibration := NewCalibration(lines[i])
			calibrations = append(calibrations, calibration)
		}

	}
	return calibrations
}

func AssessCalibrationData(input string) int {
	calibrations := parseInput(input)
	output := 0
	for i := range calibrations {
		correct := calibrations[i].Assess()
		if correct {
			output += calibrations[i].answer
		}
	}
	return output
}

func AssessCalibrationData2(input string) int {
	calibrations := parseInput(input)
	output := 0
	for i := range calibrations {
		correct := calibrations[i].Assess2()
		if correct {
			output += calibrations[i].answer
		}
	}
	return output
}

func main() {

	input, _ := os.Open("input")
	data, _ := io.ReadAll(input)
	result := AssessCalibrationData(string(data))
	result2 := AssessCalibrationData2(string(data))

	fmt.Println(result)
	fmt.Println(result2)

}
