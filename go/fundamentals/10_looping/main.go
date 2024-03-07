// iterate over the items in a collection

package main

import "fmt"

func main() {
	// for: basic
	// initialization; condition; post statement
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// for: while
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// for: infinite
	p, circuitBreaker := 0, false
	for {
		if p == 5 {
			circuitBreaker = true
		}
		if circuitBreaker {
			fmt.Println(p, circuitBreaker)
			break
		}

		p++
	}

	// continue
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
