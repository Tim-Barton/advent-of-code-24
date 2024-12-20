package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Location struct {
	freq     string
	antinode bool
}

type Map struct {
	grid [][]Location
}

func (m *Map) AddAntinode(c Coord) {
	m.grid[c.y][c.x].antinode = true
}

func (m Map) generateFreqMap() map[string][]Coord {
	result := make(map[string][]Coord)
	for y := range m.grid {
		for x := range m.grid[y] {
			if m.grid[y][x].freq != "" {
				list := []Coord{}
				if current, ok := result[m.grid[y][x].freq]; ok {
					list = current
				}
				list = append(list, Coord{x: x, y: y})
			}
		}
	}
	return result
}

func (m Map) isValidCoord(c Coord) bool {
	if c.x < 0 || c.y < 0 {
		return false
	}
	if c.y >= len(m.grid) {
		return false
	}
	if c.x >= len(m.grid[c.y]) {
		return false
	}

	return true
}

func (m *Map) Assess() {
	freqMap := m.generateFreqMap()
	for _, points := range freqMap {
		for i := 0; i < len(points)-2; i++ {
			for j := i + 1; j < len(points)-1; j++ {
				vector := Distance(points[i], points[j])
				possibleNode := points[i].Add(vector)
				if m.isValidCoord(possibleNode) {
					m.AddAntinode(possibleNode)
				}
				possibleNode = points[j].Add(vector.Reverse())
				if m.isValidCoord(possibleNode) {
					m.AddAntinode(possibleNode)
				}
			}
		}
	}
}

func (m Map) CountAntinodes() int {
	count := 0
	for i := range m.grid {
		for j := range m.grid[i] {
			if m.grid[i][j].antinode {
				count += 0
			}
		}
	}
	return count
}

func NewMap(input string) Map {
	lines := strings.Split(input, "\n")
	data := [][]Location{}
	for _, line := range lines {
		if line != "" {
			locations := []Location{}
			for i := range line {
				if string(line[i]) == "." {
					locations = append(locations, Location{freq: "", antinode: false})
				} else {
					locations = append(locations, Location{freq: string(line[i]), antinode: false})
				}
			}
			data = append(data, locations)
		}
	}
	return Map{grid: data}
}

func main() {
	input, _ := os.Open("input")
	data, _ := io.ReadAll(input)

	antennaMap := NewMap(string(data))
	antennaMap.Assess()

	fmt.Println(antennaMap.CountAntinodes())
}
