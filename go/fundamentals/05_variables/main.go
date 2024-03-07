package main

import "fmt"

func main() {
	// single creation
	var s1 = 1
	fmt.Println(s1)

	var s2 bool = true
	fmt.Println(s2)

	var s3 string
	s3 = "ok"
	fmt.Println(s3)

	// compound creation
	var c1, c2, c3 = 1, true, "ok"
	fmt.Println(c1, c2, c3)

	// block creation
	var (
		b1 int    = 1
		b2 bool   = true
		b3 string = "ok"
	)
	fmt.Println(b1, b2, b3)

	// create and assign
	ca1 := 1
	ca2, ca3 := true, "ok"

	fmt.Println(ca1, ca2, ca3)

	// variable can be reassigned
	ra1 := 1
	ra1 = 2
	ra1 = 3

	ra2 := ra1
	var ra3 = ra2

	fmt.Println(ra1, ra2, ra3)

	// variable names can only be used once per scope
	// a := 1
	// var a = 5

	// comma ok idiom
	a, b := 1, 2
	c, b := 3, 4
	fmt.Println(a, b, c)

	// defaults
	// string default ""
	// number default 0
	// other default nil

	// naming
	myLongVarName := "ok" // camel case and appropriate
	fmt.Println(myLongVarName)

	// constants
	const MaxSpeed = 100
}
