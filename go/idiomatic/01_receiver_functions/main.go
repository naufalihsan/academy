// modified function signature which allow dot notation
// allow simple mutation of existing structure
// similar to modifying a class variable in other languages

// pointer receiver can modify struct (common)
// value receiver cannot modify struct

package main

import "fmt"

type Coordinate struct {
	X, Y int
}

// regular function
func shiftBy(x, y int, c *Coordinate) {
	c.X += x
	c.Y += y
}

// receiver function (pointer)
func (c *Coordinate) shiftBy(x, y int) {
	c.X += x
	c.Y += y
}

// receiver function (value) - use copy
func (c Coordinate) shiftDist(o Coordinate) Coordinate {
	return Coordinate{X: o.X - c.X, Y: o.Y - c.Y}
}

func main() {
	coord := Coordinate{X: 1, Y: 1}
	shiftBy(1, 1, &coord)
	fmt.Println(coord)

	coord.shiftBy(-1, -1)
	fmt.Println(coord)

	point1 := Coordinate{X: 0, Y: 0}
	point2 := Coordinate{X: 2, Y: 3}
	distance := point1.shiftDist(point2)
	fmt.Println(distance)
}
