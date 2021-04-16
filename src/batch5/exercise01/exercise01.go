/*
	Hands-on exercise #1

	Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
		- first name;
		- last name;
		- favorite ice cream flavors.

	Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the
	favorite flavors.

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
		lastName:         "Lobo",
		favoriteIceCream: []string{"pequi", "tamarind", "vanilla"},
	}
	fmt.Println(p1.firstName)
	fmt.Println(p1.lastName)
	for _, v := range p1.favoriteIceCream {
		fmt.Println(v)
	}

	fmt.Println(p2.firstName)
	fmt.Println(p2.lastName)
	for _, v := range p2.favoriteIceCream {
		fmt.Println(v)
	}
}
