package main

type Coord struct {
	x int
	y int
}

func Distance(c1, c2 Coord) Vector {
	return Vector{
		x: c1.x - c2.x,
		y: c1.y - c2.y,
	}
}

func (c Coord) Add(v Vector) Coord {
	return Coord{
		x: c.x + v.x,
		y: c.y + v.y,
	}
}

type Vector struct {
	x int
	y int
}

func (v Vector) Reverse() Vector {
	return Vector{x: v.x * -1, y: v.y * -1}
}
