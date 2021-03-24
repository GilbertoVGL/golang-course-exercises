/*

	Hands-on exercise #7

	Building on the previous hands-on exercise, create a program that uses “else if” and “else”.

*/

package main

import (
	"fmt"
)

func main() {
	goodProgrammer := "I have absolutely zero idea of what I'm doing"

	if goodProgrammer == "yes you are" {
		fmt.Println("Who's a good programmer?", goodProgrammer)
	} else if goodProgrammer == "no you aren't" {
		fmt.Println("Who's a good programmer?", goodProgrammer)
	} else {
		fmt.Println("¯\\_(ツ)_/¯ ", goodProgrammer)
	}
}
