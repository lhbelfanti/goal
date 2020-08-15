package main

import (
	"fmt"
	"gitlab.com/lhbelfanti/goal/level13/exercise2/quote"
	"gitlab.com/lhbelfanti/goal/level13/exercise2/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
