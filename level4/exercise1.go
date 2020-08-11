package main

import "fmt"

func main () {
	values := [5]int{1, 2, 3, 4, 5}

	for i, v := range values {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", values)
}
