package main

import "fmt"

type person1 struct {
	firstName string
	lastName string
	favIceCreamFlavours []string
}

func main () {
	rick := person1 {
		firstName: "Rick",
		lastName: "Sanchez",
		favIceCreamFlavours: []string{"Caramel Portal", "Green apple popping isotope candy"},
	}
	morty := person1 {
		firstName: "Mortimer",
		lastName: "Smith",
		favIceCreamFlavours: []string{"Dark Matter Brownie Batter", "Galaxy Sparkle Sugar"},
	}

	m := map[string]person1{}
	m[rick.lastName] = rick
	m[morty.lastName] = morty


	for k, v := range m {
		fmt.Println("Key: ", k)
		fmt.Println("First Name: ", v.firstName)
		fmt.Println("Last Name: ", v.lastName)
		for _, f := range v.favIceCreamFlavours {
			fmt.Printf(" %v |", f)
		}
	}
}
