/*

	Hands-on exercise #10

	Using the code from the previous example, delete a record from your map. Now print the map out using the “range” loop

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
		"Ana_Fulana":    {"Sorvete", "Pastel", "Garapa"},
	}
	fmt.Println("\nBefore: ")

	for k, v := range m {
		fmt.Printf("%v:\n\t%v\n", k, v)
	}
	delete(m, "Ana_Fulana")

	fmt.Println("\nAfter: ")
	for k, v := range m {
		fmt.Printf("%v:\n\t%v\n", k, v)
	}
}
