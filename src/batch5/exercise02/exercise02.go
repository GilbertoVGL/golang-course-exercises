/*
	Hands-on exercise #2

	Take the code from the previous exercise, then store the values of type
	person in a map with the key of last name. Access each value in the map.
	Print out the values, ranging over the slice.

*/

package main

import "fmt"

type person struct {
	firstName        string
	lastName         string
	favoriteIceCream []string
}

func main() {
	p1 := person{
		firstName:        "Gilberto",
		lastName:         "Lobo",
		favoriteIceCream: []string{"soursop", "chocolate", "strawberry"},
	}
	p2 := person{
		firstName:        "José",
		lastName:         "Anhanguera",
		favoriteIceCream: []string{"pequi", "tamarind", "vanilla"},
	}

	mapping := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for k, v := range mapping {
		fmt.Printf("%v: %v\n", k, v)
		for _, iceCream := range v.favoriteIceCream {
			fmt.Printf("Favorite icecream: %v\n", iceCream)
		}
	}
}
