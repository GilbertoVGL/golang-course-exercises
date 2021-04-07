/*

	Hands-on exercise #3

	Using the code from the previous example, use SLICING to create the following new slices which are then printed:
		[42 43 44 45 46]
		[47 48 49 50 51]
		[44 45 46 47 48]
		[43 44 45 46 47]

		[0 1 2 3 4]

*/

package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := a[0:5]
	c := a[5:10]
	d := a[2:7]
	e := a[1:6]
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("d:", d)
	fmt.Println("e:", e)
}
