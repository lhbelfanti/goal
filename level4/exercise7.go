package main

import "fmt"

func main () {
	x1 := []string{ "James", "Bond", "Shaken, not stirred"}
	x2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	x := [][]string {x1, x2}

	for i, v := range x {
		fmt.Println("record num: ", i)
		for j, v2 := range v {
			fmt.Printf("\tindex: %v | value: %v\n", j, v2)
		}
	}
}
