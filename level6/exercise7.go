package main

import "fmt"

var g func(s string)

func main () {
	sh := "Sherlock Holmes"
	bin := 	func(s string) {
		fmt.Println(s, "in binary is:")
		for _, v := range s{
			fmt.Printf("%b", v)
		}
	}
	bin(sh)

	fmt.Printf("\n%T\n", bin)

	g = bin
	g(sh)
	fmt.Printf("\nThis is g %T\n", g)
}
