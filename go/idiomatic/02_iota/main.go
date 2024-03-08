// common to make groups of constants
// iota (integer to constants)
// iota keywords can be used to automatically assign values

package main

import "fmt"

// const (
// 	Online      = 1
// 	Offline     = 2
// 	Maintenance = 3
// 	Retired     = 4
// )

// long form
// const (
// 	Online      = iota
// 	Offline     = iota
// 	Maintenance = iota
// 	Retired     = iota
// )

// short form
const (
	Online = iota
	Offline
	Maintenance
	Retired
)

// skip a value
const (
	s0 = iota
	_
	_
	s3
	s4
)

// start at
const (
	i3 = iota + 3
	i4
	i5
)

type Direction byte

const (
	North Direction = iota
	East
	South
	West
)

// iota enumeration pattern
func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

func (d Direction) Move() string {
	switch d {
	case North:
		return "Top"
	case East:
		return "Right"
	case South:
		return "Bottom"
	case West:
		return "Left"
	default:
		return "Center"
	}
}

func main() {
	d := North
	fmt.Println(d.String())
	fmt.Println(d.Move())
}
