// https://goosegeesejeez.gitbook.io/apprendre-go-par-les-tests/fondamentaux-de-go/hello-world
package main

const (
	english = "English"
	french  = "French"
	spanish = "Spanish"
)

const (
	englishHelloPrefix = "Hello, "
	frenchHelloPrefix  = "Bonjour, "
	spanishHelloPrefix = "Hola, "
)

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return // implicit return prefix
}

func Bonjour(name string, lang string) string {
	if name == "" {
		name = "Monde"
	}

	return greetingPrefix(lang) + name
}
