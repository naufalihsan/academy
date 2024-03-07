// program code executes line by line
// flow control is a way to interrupt this process

package main

import (
	"errors"
	"fmt"
)

func getApproval(age int) bool {
	return age > 21
}

func getPosition(age int) string {
	if age > 21 {
		return "manager"
	} else if age > 18 && age <= 21 {
		return "freshman"
	} else {
		return "student"
	}
}

func getClearence(age int) (bool, error) {
	if age <= 21 {
		return false, errors.New("restricted")
	}
	return true, nil
}

func main() {
	age := 22
	parentalApproval, govermentApproval := false, false

	// if .. else
	if age > 21 {
		fmt.Println("ok")
	}

	if age > 21 {
		fmt.Println("ok")
	} else {
		fmt.Println("restricted")
	}

	// if .. else if
	if age > 21 {
		fmt.Println("ok")
	} else if parentalApproval {
		fmt.Println("parental advisory")
	} else {
		fmt.Println("restricted")
	}

	// logical operators
	if parentalApproval && govermentApproval {
		fmt.Println("ok")
	}

	// usage with functions
	if getApproval(age) {
		fmt.Println("ok")
	}

	// statement initialization
	if rank := getPosition(age); rank == "manager" {
		fmt.Println("ok")
	} else if rank == "freshman" {
		fmt.Println("parental advisory")
	} else {
		fmt.Println("restricted")
	}

	// early return
	clearence, err := getClearence(age)
	if err != nil {
		return
	}
	fmt.Println(clearence)
}
