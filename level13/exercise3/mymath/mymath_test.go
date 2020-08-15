package mymath

import (
	"fmt"
	"testing"
)

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{1, 2, 3, 4, 3, 2, 1})
	}
}

func ExampleCenteredAvg() {
	a := []int{1, 4, 6, 8, 100}
	fmt.Println(CenteredAvg(a))
	// Output:
	// 6
}

func TestCenteredAvg(t *testing.T) {
	type data struct {
		n []int
		r float64
	}

	a := []int{1, 4, 6, 8, 100}
	b := []int{0, 8, 10, 1000}
	c := []int{9000, 4, 10, 8, 6, 12}
	d := []int{123, 744, 140, 200}

	x := []data {
		{n: a, r: 6},
		{n: b, r: 9},
		{n: c, r: 9},
		{n: d, r: 170},
	}

	for _, v := range x {
		if f := CenteredAvg(v.n); f != v.r {
			t.Error("Expected:", v.r, "got:", f)
		}
	}
}
