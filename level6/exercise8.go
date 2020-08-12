package main

import "fmt"

func main () {
	x := toBinary("John Watson")
	x()
}

func toBinary(s string) func() {
	return func() {
		fmt.Println(s, "in binary is:")
		for _, v := range s{
			fmt.Printf("%b", v)
		}
	}
}
