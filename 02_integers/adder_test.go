// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/integers
package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

// This is a example invoking the `Add` func.
// It should start with the `Example` word.
// The last commant line is mandatory for the example to be caugth by the `go test -v` command.
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestRepeatBetter(t *testing.T) {
	repeated := RepeatBetter("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a")
	}
}

func BenchmarkRepeatBetter(b *testing.B) {
	for b.Loop() {
		Repeat("a")
	}
}
