// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/integers
package integers

import "strings"

// Add takes two integers and returns the sum of them.
func Add(x, y int) int {
	return x + y
}

func Repeat(letter string) (repeated string) {
	for range 5 {
		repeated += letter
	}
	return
}

func RepeatBetter(letter string, repeat int) string {
	var repeated strings.Builder

	for range repeat {
		repeated.WriteString(letter)
	}

	return repeated.String()
}
