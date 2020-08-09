package main

import "fmt"

type apple1 int
var x3 apple1
var y2 int

func main() {
	fmt.Println(x3)
	fmt.Printf("%T\n", x3)
	x3 = 42
	fmt.Println(x3)
	y2 = int(x3)
	fmt.Println(y2)
	fmt.Printf("%T", y2)
}
