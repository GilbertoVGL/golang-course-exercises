/*

	Hands-on exercise #4

	FYI - nice documentation and new terminology “underlying type”: https://golang.org/ref/spec#Types

	For this exercise:
		- Create your own type. Have the underlying type be an int.
		- Create a VARIABLE of your new TYPE with the IDENTIFIER “x” using the “VAR” keyword

	in func main
		- Print out the value of the variable “x”
		- Print out the type of the variable “x”
		- Assign 42 to the VARIABLE “x” using the “=” OPERATOR
		- Print out the value of the variable “x”

*/

package main

import "fmt"

type gopher int

var x gopher

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42

	fmt.Println(x)
}
