package main

import "fmt"

func main () {
	x := foo()
	fmt.Println(x)
	y, z := bar()
	fmt.Println(y, z)
}

func foo() int {
	return 27
}

func bar() (int, string) {
	return 27, "Holmes"
}
