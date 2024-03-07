// go is statically typed language (must be provided)
// go use type inference to determine what type of data it is working with (compile error if wrong type used)

package main

import "fmt"

func main() {
	// all primitive data types are numeric in Go

	// signed integer types
	// int8		signed  8-bit integers (-128 to 127)
	// int16    signed 16-bit integers (-32768 to 32767)
	// int32    signed 32-bit integers (-2147483648 to 2147483647)
	// int64    signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

	var minInt8 int8 = -128
	var maxInt8 int8 = 127
	fmt.Printf("int8 min: %d - max: %d\n", minInt8, maxInt8)

	// unsigned integer types
	// uint8     unsigned  8-bit integers (0 to 255)
	// uint16    unsigned 16-bit integers (0 to 65535)
	// uint32    unsigned 32-bit integers (0 to 4294967295)
	// uint64    unsigned 64-bit integers (0 to 18446744073709551615)

	var minUInt8 uint8 = 0
	var maxUInt8 uint8 = 255
	fmt.Printf("uint8 min: %d - max: %d\n", minUInt8, maxUInt8)

	// other data types
	// float32     IEEE-754 32-bit floating-point numbers
	// float64     IEEE-754 64-bit floating-point numbers
	// complex64   complex numbers with float32 real and imaginary parts
	// complex128  complex numbers with float64 real and imaginary parts
	// byte        alias for uint8
	// rune        alias for int32
	// bool 	   true or false

	// Type Aliases
	type Speed int
	type Velocity Speed

	var maxSpeed Speed = Speed(100)
	var maxVelocity Velocity = Velocity(maxSpeed)

	fmt.Printf("Max velocity: %d\n", maxVelocity)
}
