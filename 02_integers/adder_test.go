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
