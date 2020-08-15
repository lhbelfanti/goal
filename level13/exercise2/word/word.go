// Package word is used to count words from a string
package word

import (
	"strings"
)

// UseCount returns the number of times each word is used in string
// no need to write an example for this one
// writing a test for this one is a bonus challenge; harder
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count returns the number of words in a string
func Count(s string) int {
	return len(strings.Fields(s))
}
