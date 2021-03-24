/*

	Hands-on exercise #8

	Create a program that uses a switch statement with no switch expression specified.

*/

package main

import (
	"fmt"
)

func main() {
	var typedSwitchTest interface{}
	typedSwitchTest = getSomeValue()

	switch t := typedSwitchTest.(type) {
	case int:
		fmt.Println("You are int")
	case bool:
		fmt.Println("You are bool")
	case string:
		fmt.Println("You are string")
	case float64:
		fmt.Println("You are float")
	default:
		fmt.Printf("WTH are you? %T\n", t)
	}
}

func getSomeValue() int {
	x := int(1)
	return x
}
