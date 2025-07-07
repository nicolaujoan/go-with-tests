package main

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "

	defaultName = "World"
)

func Hello(name string, language string) string {
	if name == "" {
		name = defaultName
	}

	return greetingPrefix(language) + name
}

// lowercase because it is a private function
func greetingPrefix(language string) (prefix string) { // named return value
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Chris", ""))
}