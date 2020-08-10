package main

import "fmt"

func main () {
	i := 2
	j := 3

	if i == j { // false
		fmt.Println("i == j")
	}

	if i != j { // true
		fmt.Println("i != j")
	}
}
