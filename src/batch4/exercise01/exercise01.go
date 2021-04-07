/*
	Hands-on exercise #1

	Using a COMPOSITE LITERAL:
		- Create an ARRAY which holds 5 VALUES of TYPE int
		- Assign VALUES to each index position.
	Range over the array and print the values out.
	Using format printing:
		- print out the TYPE of the array


*/

package main

import "fmt"

func main() {
	myArray := [5]int{0, 1, 2, 3, 4}
	for i, v := range myArray {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", myArray)
}
