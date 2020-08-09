package main

import "fmt"

type apple int
var x2 apple

func main() {
	fmt.Println(x2)
	fmt.Printf("%T\n", x2)
	x2 = 42
	fmt.Println(x2)
}
