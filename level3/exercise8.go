package main

import "fmt"

func main () {

	switch {
	case (2 == 3):
			fmt.Println("should not print")
	case (3 == 3):
		fmt.Printf("should print")
	}

}
