/*

	Hands-on exercise #8

	Create a map with a key of TYPE string which is a person’s “last_first” name,
	and a value of TYPE []string which stores their favorite things. Store three records
	in your map. Print out all of the values, along with their index position in the slice.

	`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
	`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
	`no_dr`, `Being evil`, `Ice cream`, `Sunsets`

*/

package main

import (
	"fmt"
)

func main() {
	m := map[string][]string{
		"Gilberto_Lobo": {"Play games", "Learn", "Sports"},
		"Jose_Alguem":   {"Correr", "Nadar", "Comer"},
		"Jose_Ninguem":  {"Alpinismo", "Plantas", "Bicicleta"},
	}

	for k, v := range m {
		fmt.Printf("%v\t", k)
		for _, vv := range v {
			fmt.Printf("%v\t", vv)
		}
		fmt.Printf("\n")
	}
}
