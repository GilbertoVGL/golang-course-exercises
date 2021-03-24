/*

	Hands-on exercise #9

	Create a program that uses a switch statement with the switch expression specified
	as a variable of TYPE string with the IDENTIFIER “favSport”.

*/

package main

import (
	"fmt"
)

func main() {
	favSport := "Dota2"
	switch favSport {
	case "LOL":
		fmt.Println("No way man")
	case "Dota2":
		fmt.Println("Sounds good ma")
	}
}
