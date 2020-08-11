package main

import "fmt"

func main () {
	values := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(values[:5])
	fmt.Println(values[5:])
	fmt.Println(values[2:7])
	fmt.Println(values[1:6])
}
