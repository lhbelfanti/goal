package main

import "fmt"

func main () {
	const (
		typed int = 1
		untyped = 2
	)

	fmt.Println(typed)
	fmt.Println(untyped)
}
