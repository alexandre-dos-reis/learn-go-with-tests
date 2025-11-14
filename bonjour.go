// https://goosegeesejeez.gitbook.io/apprendre-go-par-les-tests/fondamentaux-de-go/hello-world
package main

import "fmt"

func Bonjour(name string) string {
	return "Bonjour, " + name
}

func main() {
	fmt.Println(Bonjour("monde"))
}
