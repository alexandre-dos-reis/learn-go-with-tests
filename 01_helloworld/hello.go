// https://goosegeesejeez.gitbook.io/apprendre-go-par-les-tests/fondamentaux-de-go/hello-world
package hello

// NOTE: Go source files can only have one package per directory. Make sure that your files are organised into their own packages.

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
