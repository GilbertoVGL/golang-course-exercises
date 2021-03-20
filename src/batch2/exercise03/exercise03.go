/*
	Hands-on exercise #3

	Create TYPED and UNTYPED constants. Print the values of the constants.

*/

package main

import "fmt"

const (
	typedConst   string = "coool"
	untypedConst        = "also cool"
)

func main() {
	fmt.Println(typedConst)
	fmt.Println(untypedConst)
}
