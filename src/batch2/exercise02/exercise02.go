/*
	Hands-on exercise #2

	Using the following operators, write expressions and assign their values to variables:
		==
		<=
		>=
		!=
		<
		>

	Now print each of the variables.

*/

package main

import "fmt"

func main() {
	a := 1991 == 2021
	b := 1991 <= 2021
	c := 1991 >= 2021
	d := 1991 != 2021
	e := 1991 < 2021
	f := 1991 > 2021

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
