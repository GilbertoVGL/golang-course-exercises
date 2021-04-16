/*

	Hands-on exercise #3

		- Create a new type called vehicle.
			- The underlying type is a struct;
			- The fields:
				doors;
				color.

		- Create two new types: truck & sedan.
			- The underlying type of each of these new types is a struct;
			- Embed the “vehicle” type in both truck & sedan;
			- Give truck the field “fourWheel” which will be set to bool;
			- Give sedan the field “luxury” which will be set to bool. solution.

		- Using the vehicle, truck, and sedan structs:
			- using a composite literal, create a value of type truck and assign values to the fields;
			- using a composite literal, create a value of type sedan and assign values to the fields.

		- Print out each of these values.
		- Print out a single field from each of these values.


*/

package main

import "fmt"

type vehicle struct {
	doors int
	color int16
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: 0xfa,
		},
		fourWheel: true,
	}
	s1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: 0x0,
		},
		luxury: true,
	}

	fmt.Printf("%+v\n", t1)
	fmt.Printf("%+v\n", t1.doors)
	fmt.Printf("%+v\n", s1)
	fmt.Printf("%+v\n", s1.color)
}
