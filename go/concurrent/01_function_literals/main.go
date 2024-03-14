// define a function within a function
// also known as closures or anonymous functions

package main

import (
	"fmt"
	"strings"
)

// anonymous function
func start() {
	fmt.Println("Hello")
	world := func() {
		fmt.Println("World")
	}
	world()
}

// function params
func customMsg(fn func(m string), m string) {
	fn(m)
}

func header() func(m string) {
	return func(m string) {
		border := strings.Repeat("-", len(m))

		fmt.Printf("%.*s\n", len(m), border)
		fmt.Println(m)
		fmt.Printf("%.*s\n", len(m), border)
	}
}

// type aliases are helpful when passing function literals to other functions
type DiscountFn func(subTotal float64) float64

func calculatePrice(subTotal float64, discountFn DiscountFn) float64 {
	return subTotal - (subTotal * discountFn(subTotal))
}

func main() {
	start()

	title := "Hello World"
	customMsg(header(), title)

	// closure (capture surrounding params)
	discount := 0.2
	discountFn := func(subTotal float64) float64 {
		if subTotal > 100.0 {
			discount += 0.1
		}

		if discount > 0.3 {
			discount = 0.3
		}
		return discount
	}
	totalPrice := calculatePrice(100.0, discountFn)
	fmt.Println(totalPrice)
}
