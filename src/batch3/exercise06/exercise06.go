/*

	Hands-on exercise #6

	Create a program that shows the “if statement” in action.

*/

package main

import (
	"fmt"
)

func main() {
	goodProgrammer := false
	if goodProgrammer {
		fmt.Println("It's", goodProgrammer, "that you are a good programmer")
	} else {
		fmt.Println("Sorry sir, but you are not allowed in this club.")
	}
}
