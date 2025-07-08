package iteration

import (
	"testing"
	"fmt"
	"strings"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 2)
	expected := "aa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestRepeatStdLib(t *testing.T) {
	repeated := strings.Repeat("a", 2)
	expected := "aa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

// Benchmarks are first class citizens
// go test -bench=.
func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 2)
	}
}

func BenchmarkRepeatStdLib(b *testing.B) {
	for b.Loop() {
		strings.Repeat("a", 2)
	}
}

// the output comment is needed to execute the example
func ExampleRepeat() {
	repeated := Repeat("a", 2)
	fmt.Println(repeated)
	// Output: aa
}