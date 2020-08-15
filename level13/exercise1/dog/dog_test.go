package dog

import (
	"fmt"
	"testing"
)

type data struct {
	hy int
	dy int
}

func TestYears(t *testing.T) {
	values := []data {
		{hy: 1, dy: 7},
		{hy: 7, dy: 49},
	}

	for _, v := range values {
		if years := Years(v.hy); years != v.dy {
			t.Error("Expected", v.dy, "Got", years)
		}
	}
}

func ExampleYears() {
	fmt.Println(Years(7))
	// Output:
	// 49
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(7)
	}
}

func TestYears2(t *testing.T) {
	values := []data {
		{hy: 1, dy: 7},
		{hy: 7, dy: 49},
	}

	for _, v := range values {
		if dy := YearsTwo(v.hy); dy != v.dy {
			t.Error("Expected", v.dy, "Got", dy)
		}
	}
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(7))
	// Output:
	// 49
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(7)
	}
}
