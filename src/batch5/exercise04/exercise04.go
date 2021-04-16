/*

	Hands-on exercise #4

	Create and use an anonymous struct.

*/

package main

import "fmt"

func main() {
	x := struct {
		game     string
		platform []string
		size     int
		rating   int
		tag      int
	}{
		game:     "Valheim",
		platform: []string{"PC", "PlayStation 4", "PlayStation 5", "XBox One"},
		size:     508,
		rating:   5,
		tag:      0,
	}

	fmt.Printf("%+v\n", x)
}
