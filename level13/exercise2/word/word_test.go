package word

import (
	"fmt"
	"gitlab.com/lhbelfanti/goal/level13/exercise2/quote"
	"reflect"
	"testing"
)

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func ExampleCount() {
	fmt.Println(Count("Words are great"))
	// Output:
	// 3
}

func TestCount(t *testing.T) {
	r := Count("Words are great")
	if r != 3 {
		t.Error("Expected:", 3, "got:", r)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}

func ExampleUseCount() {
	fmt.Println(UseCount("Colors blue yellow blue"))
	// Output:
	// map[Colors:1 blue:2 yellow:1]
}

func TestUseCount(t *testing.T) {
	r := UseCount("Colors blue yellow blue")
	e := map[string] int{"Colors":1, "blue":2, "yellow":1}

	eq := reflect.DeepEqual(r, e)
	if !eq {
		t.Error("Expected:", r, "got:", e)
	}
}




