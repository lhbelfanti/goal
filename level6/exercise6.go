package main

import "fmt"

func main () {
	func() {
		sh := "Sherlock Holmes"
		for _, v := range sh{
			fmt.Printf("%b", v)
		}
	}()
}
