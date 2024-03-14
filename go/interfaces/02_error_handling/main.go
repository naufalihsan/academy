// go has no exceptions
// error are returned as the last return value from a function

package main

import (
	"errors"
	"fmt"
)

// basic error
func divide(lhs, rhs int) (int, error) {
	if rhs == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return rhs / lhs, nil
}

// implement error interface
type DivError struct {
	a, b int
}

// always implement error as a receiver function
func (d *DivError) Error() string {
	return fmt.Sprintf("cannot divide by zero: %d / %d", d.a, d.b)
}

func division(lhs, rhs int) (int, error) {
	if rhs == 0 {
		return 0, &DivError{lhs, rhs}
	}
	return rhs / lhs, nil
}

func main() {
	a, b := 10, 0
	r, err := division(a, b)
	if err != nil {
		var divErr *DivError
		if errors.As(err, &divErr) {
			fmt.Println(err)
		}
		return
	}
	fmt.Println(r)
}
