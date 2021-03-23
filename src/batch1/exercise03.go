/*
	Hands-on exercise #3

	Using the code from the previous exercise, at the package level scope,
	assign the following values to the three variables:
		- for x assign 42
		- for y assign “James Bond”
		- for z assign true

	in func main:
		- Use fmt.Sprintf to print all of the VALUES to one single string.
		- ASSIGN the returned value of TYPE string using the short declaration operator to a VARIABLE with the IDENTIFIER “s”
		- Print out the value stored by variable “s”
*/

package main

import "fmt"

var x int
var y string
var z bool

func main() {
	x = 42
	y = "James Bond"
	z = true

	s := fmt.Sprintf("%d %s %t", x, y, z)

	fmt.Println(s)
}
