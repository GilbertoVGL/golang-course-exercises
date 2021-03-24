/*

	Hands-on exercise #4

	Create a for loop using this syntax:
		for { }

	Have it print out the years you have been alive.


*/

package main

import "fmt"

func main() {
	birthYear := 1991
	currentYear := 2021

	for {
		fmt.Println(birthYear)
		birthYear++

		if birthYear > currentYear {
			break
		}
	}
}
