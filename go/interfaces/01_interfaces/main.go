// interface are implicity implmented
// when type has all receiver functions required by the inteface then it's considered implemented
// prefer multiple interface with few functions over one large interface

package main

import "fmt"

// case 1
type MyInteface interface {
	Function1()
	Function2(x int) int
}

type MyType int

func (m MyType) Function1() {
	fmt.Println("MyType run Function1()")
}
func (m MyType) Function2(x int) int {
	return int(m) * x * 2
}

type AnotherType struct {
	number int
}

func (a *AnotherType) Function1() {
	fmt.Println("AnotherType run Function1()")
}

func (a *AnotherType) Function2(x int) int {
	return a.number * x * 2
}

func execute(i MyInteface) {
	i.Function1()

	var r int
	switch v := i.(type) {
	case MyType:
		r = v.Function2(10)
	case *AnotherType:
		r = v.Function2(20)
	}
	fmt.Println(r)
}

// case 2
type Resetter interface {
	Reset()
}

type Coordinate struct {
	X, Y int
}

type Player struct {
	health   int
	position Coordinate
}

func (p *Player) Reset() {
	p.health = 100
	p.position = Coordinate{X: 0, Y: 0}
}

func Reset(r Resetter) {
	r.Reset()
}

// access the underlying type that implements an interface
func ResetWithPenalty(r Resetter) {
	if player, ok := r.(*Player); ok {
		player.health = 50
	} else {
		r.Reset()
	}
}

func main() {
	m := MyType(1)
	execute(m)
	// execute(&m)

	a := AnotherType{number: 1}
	execute(&a)

	p := Player{health: 50, position: Coordinate{5, 5}}
	fmt.Println(p)
	Reset(&p)
	fmt.Println(p)
	ResetWithPenalty(&p)
	fmt.Println(p)
}
