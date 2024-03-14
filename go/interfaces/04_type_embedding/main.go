// provide existing functionality to new type
// require a type to implement multiple interfaces

package main

import "fmt"

// embedded interface
type Whisperer interface {
	Whisper() string
}

type Yeller interface {
	Yell() string
}

type Talker interface {
	Whisperer
	Yeller
}

func Talk(t Talker) {
	fmt.Println(t.Whisper())
	fmt.Println(t.Yell())
}

// embedded structs
type Collar struct {
	vendor string
	size   int
}

type Dog struct {
	name string
	Collar
}

func (d Dog) Whisper() string {
	return "🐕"
}

func (d Dog) Yell() string {
	return "🐶"
}

func (c Collar) GetColar() {
	fmt.Printf("Vendor: %v Size: %d\n", c.vendor, c.size)
}

func main() {
	dog := Dog{"Scooby", Collar{vendor: "iPaw", size: 8}}
	dog.GetColar()

	Talk(dog)
}
