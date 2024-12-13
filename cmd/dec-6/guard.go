package main

import "fmt"


type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Guard struct {
	x           int
	y           int
	direction   Direction
	stepHistory map[string]interface{}
}

func (g *Guard) turn() {
	switch g.direction {
	case Up:
		g.direction = Right
	case Right:
		g.direction = Down
	case Down:
		g.direction = Left
	case Left:
		g.direction = Up
	}
}

func (g Guard) Steps() int {
	return len(g.stepHistory)
}

func (g *Guard) PatrolStep(l Lab) bool {
	//fmt.Printf("Tick. Guard(%d,%d)\n", g.x, g.y)
	for {

		x, y := g.x, g.y
		switch g.direction {
		case Up:
			y = y - 1
		case Down:
			y = y + 1
		case Left:
			x = x - 1
		case Right:
			x = x + 1
		}

		if x < 0 || x > len(l.grid[0])-1 || y < 0 || y > len(l.grid)-1 {
			return false
		}
		//fmt.Printf("X Length: %d, Y Length: %d\n", len(l.grid[y]), len(l.grid))
		//fmt.Printf("X: %d, Y:%d\n", x, y)
		aheadSpace := l.grid[y][x]

		switch aheadSpace {
		case Empty:
			g.x = x
			g.y = y
			return true
		case Obstructed:
			g.turn()
		}
	}
}

func (g *Guard) Patrol(lab Lab) {
	guardMoving := true
	for guardMoving {
		guardMoving = g.PatrolStep(lab)
		stepIndex := fmt.Sprintf("%d,%d", g.x, g.y)
		_, ok := g.stepHistory[stepIndex]
		if !ok {
			g.stepHistory[stepIndex] = nil
		}
	}
}
