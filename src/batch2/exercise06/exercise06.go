/*

	Hands-on exercise #5

	Using iota, create 4 constants for the NEXT 4 years. Print the constant values.

*/

package main

import (
	"fmt"
)

const (
	_     = iota
	year1 = 2021 + iota
	year2 = 2021 + iota
	year3 = 2021 + iota
	year4 = 2021 + iota
)

func main() {
	fmt.Println(year1)
	fmt.Println(year2)
	fmt.Println(year3)
	fmt.Println(year4)
}
