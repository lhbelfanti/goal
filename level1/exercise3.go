package main

import "fmt"

var x1 int = 42
var y1 string = "James Bond"
var z1 bool = true

func main() {
	s := fmt.Sprintf("%d, %s, %t", x1, y1, z1)
	//s := fmt.Sprintf("%v, %v, %v", x1, y1, z1) // -> %v is by default the value of the variable
	fmt.Println(s)
}
