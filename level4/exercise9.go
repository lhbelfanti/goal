package main

import "fmt"

func main () {
	x := map[string][]string {
		"bond_james": {`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": {`James Bond`, `Literature`, `Computer Science`},
		"no_dr": {`Being evil`, `Ice cream`, `Sunsets`},
	}

	x["rob_pike"] = []string{"programming", "go", "unix"}

	for k, s := range x {
		fmt.Println("Key: ", k)
		for i, v := range s {
			fmt.Printf("\tindex: %v | value: %v\n", i, v)
		}
	}
}
