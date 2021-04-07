/*

	Hands-on exercise #9

	Using the code from the previous example, add a record to your map. Now print the map out using the “range” loop

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

	m["Ana_Fulana"] = []string{"Sorvete", "Pastel", "Garapa"}

	for k, v := range m {
		fmt.Printf("%v:\n\t%v\n", k, v)
	}
}
