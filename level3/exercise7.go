package main

import "fmt"

func main () {
	i := 2
	j := 3
	k := 2

	if i == j { // false
		fmt.Println("i == j")
	} else if j == k { // false
		fmt.Println("j == k")
	} else if i == k { // true
		fmt.Println("i == k")
	} else {
		fmt.Printf("default option")
	}
}
