package main

import "fmt"

func main () {
	values := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for i, v := range values {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", values)
}
