package main

import "fmt"

type person struct {
	firstName string
	lastName string
	favIceCreamFlavours []string
}

func main () {
	rick := person {
		firstName: "Rick",
		lastName: "Sanchez",
		favIceCreamFlavours: []string{"Caramel Portal", "Green apple popping isotope candy"},
	}
	morty := person {
		firstName: "Mortimer",
		lastName: "Smith",
		favIceCreamFlavours: []string{"Dark Matter Brownie Batter", "Galaxy Sparkle Sugar"},
	}

	fmt.Println("First Name: ", rick.firstName)
	fmt.Println("Last Name: ", rick.lastName)
	fmt.Print("Favourite ice cream flavours: ")
	for _, v := range rick.favIceCreamFlavours {
		fmt.Printf(" %v |", v)
	}
	fmt.Println()

	fmt.Println("First Name: ", morty.firstName)
	fmt.Println("Last Name: ", morty.lastName)
	fmt.Print("Favourite ice cream flavours: ")
	for _, v := range morty.favIceCreamFlavours {
		fmt.Printf(" %v |", v)
	}
}
