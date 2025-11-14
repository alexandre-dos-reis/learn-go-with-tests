// https://goosegeesejeez.gitbook.io/apprendre-go-par-les-tests/fondamentaux-de-go/hello-world
package main

import "fmt"

const (
	anglais  = "Anglais"
	francais = "Français"
	espagnol = "Espagnol"
)

const (
	prefixeSalutAnglais  = "Hello, "
	prefixeSalutFrancais = "Bonjour, "
	prefixeSalutEspagnol = "Hola, "
)

func Bonjour(name string) string {
	if name == "" {
		name = "Monde"
	}

	return prefixeSalutFrancais + name
}

func main() {
	fmt.Println(Bonjour("monde"))
}
