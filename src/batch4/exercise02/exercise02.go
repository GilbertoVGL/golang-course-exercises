/*
	Hands-on exercise #2

	Using a COMPOSITE LITERAL:
		- create a SLICE of TYPE int
		- assign 10 VALUES
	Range over the slice and print the values out.
	Using format printing:
		- print out the TYPE of the slice

*/

package main

import "fmt"

func main() {
	mySlice := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}

	for i, v := range mySlice {
		fmt.Println((i + 1), v)
	}

	fmt.Printf("%T", mySlice)
}
