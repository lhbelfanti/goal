package main

import "fmt"

func main () {
	for i := 10; i < 101; i++ {
		fmt.Printf("%d divided by 4 -> reminder = %d \n", i, i%4)
	}
}
