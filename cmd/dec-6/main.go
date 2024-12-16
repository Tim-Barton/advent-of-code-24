package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

type PositionType int

const (
	Empty PositionType = iota
	Obstructed
)

var stateName = map[PositionType]string{
	Empty:      ".",
	Obstructed: "#",
}

func (pt PositionType) String() string {
	return stateName[pt]
}

type Lab struct {
	grid [][]PositionType
}

func (l Lab) Copy() Lab {
	newGrid := [][]PositionType{}
	for i := range l.grid {
		newGridLine := []PositionType{}
		for j := range l.grid[i] {
			newGridLine = append(newGridLine, l.grid[i][j])
		}
		newGrid = append(newGrid, newGridLine)
	}
	return Lab{grid: newGrid}
}

func (l Lab) String() string {
	output := ""
	for i := range l.grid {
		for j := range l.grid[i] {
			//if l.guard.x == j && l.guard.y == i {
			//output += "^"
			//} else {
			output += l.grid[i][j].String()
			//}
		}
		output += "\n"
	}
	return output
}

func parseLine(input string, line int) ([]PositionType, *Guard) {
	output := []PositionType{}
	var guard *Guard
	for i := range input {
		switch input[i] {
		case '.':
			output = append(output, Empty)
		case '#':
			output = append(output, Obstructed)
		default:
			output = append(output, Empty)
			guard = &Guard{x: i, y: line, direction: Up, stepHistory: make(map[string]interface{})}
		}
	}
	return output, guard
}

func parseInput(s string) (Lab, Guard) {
	lab := Lab{}
	outputGuard := Guard{}
	data := [][]PositionType{}
	lineCount := 0
	lines := strings.Split(s, "\n")
	//fmt.Printf("Loaded lines from input: %d\n", len(lines))
	for i := range lines {
		line := lines[i]
		if line == "" {
			continue
		}
		localData, guard := parseLine(line, lineCount)
		if guard != nil {
			outputGuard = *guard
		}
		data = append(data, localData)
		lineCount += 1
		//fmt.Printf("Line: %d, Length: %d\n", lineCount, len(data))
	}
	lab.grid = data
	return lab, outputGuard
}

func loopDetector(g Guard, l Lab) bool {
	doubleStepper := NewDoubleStepper(g)
	continuePatrol := true
	for continuePatrol {
		g.PatrolStep(l)
		continuePatrol = doubleStepper.PatrolStep(l)
		if Equals(g, doubleStepper) {
			return true
		}
	}
	return false
}

func FindObstructionPoints(g Guard, l Lab) int {

	collate := make(chan bool, 100)
	jobs := make(chan Lab, 100)
	jobsGroup := sync.WaitGroup{}
	count := 0
	jobsGroup.Add(1)
	go func() {
		for i := range l.grid {
			for j := range l.grid[i] {
				if l.grid[i][j] == Empty {
					jobsGroup.Add(1)
					newLab := l.Copy()
					newLab.grid[i][j] = Obstructed
					jobs <- newLab
				}
			}
		}
		jobsGroup.Done()
	}()

	go func() {
		count := 0
		for job := range jobs {
			collate <- loopDetector(g, job)
			count += 1
		}
	}()

	go func() {
		for result := range collate {
			if result {
				count += 1
			}
			jobsGroup.Done()
		}
	}()

	jobsGroup.Wait()
	close(jobs)
	close(collate)

	return count

}

func main() {

	input, _ := os.Open("input")
	data, _ := io.ReadAll(input)

	lab, guard := parseInput(string(data))
	guard.Patrol(lab)

	fmt.Println(guard.Steps())

	lab, guard = parseInput(string(data))
	count := FindObstructionPoints(guard, lab)

	fmt.Println(count)
}
