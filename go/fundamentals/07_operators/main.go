// operators used to perform calculation and comparisons
// work on operands

package main

import "fmt"

func main() {
	// arithmetic operators (always return a number)
	// 	+	adds two operands
	// 	-	subtracts second operand from the first
	// 	*	multiplies both operands
	// 	/	divides the numerator by the denominator.
	// 	%	modulus operator; gives the remainder after an integer division.
	// 	++	increment operator. it increases the integer value by one.
	// 	--	decrement operator. it decreases the integer value by one.

	var s = 3 + 3
	fmt.Println(s)

	a := 1
	a += 3 // equals a = a + 3
	a++
	fmt.Println(a)

	// relational operators (always return a boolean)
	// ==	if the values of two operands are equal or not; if yes, the condition becomes true.
	// !=	if the values of two operands are equal or not; if the values are not equal, then the condition becomes true.
	// >	if the value of left operand is greater than the value of right operand; if yes, the condition becomes true.
	// <	if the value of left operand is less than the value of the right operand; if yes, the condition becomes true.
	// >=	if the value of the left operand is greater than or equal to the value of the right operand; if yes, the condition becomes true
	// <=	if the value of left operand is less than or equal to the value of right operand; if yes, the condition becomes true.

	a, b := 1, 2
	var aEqualsB = a == b
	fmt.Println(aEqualsB)

	// logical operators
	// &&	logical AND operator. If both the operands are non-zero, then condition becomes true.
	// ||	logical OR Operator. If any of the two operands is non-zero, then condition becomes true.
	// !	logical NOT Operator. Use to reverses the logical state of its operand. If a condition is true then Logical NOT operator will make false.

	c, d := true, false
	var cAndD = c && d
	fmt.Println(cAndD)

	// bitwise operators
	// &	binary AND Operator copies a bit to the result if it exists in both operands.
	// |	binary OR Operator copies a bit if it exists in either operand.
	// ^	binary XOR Operator copies the bit if it is set in one operand but not both.
	// <<	binary Left Shift Operator. The left operands value is moved left by the number of bits specified by the right operand.
	// >>	binary Right Shift Operator. The left operands value is moved right by the number of bits specified by the right operand.

	var e uint = 60 /* 60 = 0011 1100 */
	var f uint = 13 /* 13 = 0000 1101 */
	var g uint = 0

	g = e & f  /* 12 = 0000 1100 */
	g = e | f  /* 61 = 0011 1101 */
	g = e ^ f  /* 49 = 0011 0001 */
	g = e << 2 /* 240 = 1111 0000 */
	g = e >> 2 /* 15 = 0000 1111 */

	fmt.Println(g)

	// miscellaneous operators
	// &	returns the address of a variable.
	// *	pointer to a variable.

	var h int = 10
	var pointerH *int = &h

	fmt.Println(h, pointerH, *pointerH)

}
