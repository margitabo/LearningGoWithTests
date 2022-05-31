package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const germanHelloPrefix = "Hallo, "
const spanish = "Spanish"
const french = "French"
const german = "German"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case french:
		prefix = frenchHelloPrefix

	case spanish:
		prefix = spanishHelloPrefix

	case german:
		prefix = germanHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	//In our function signature we have made a named return value (prefix string).
	//You can return whatever it's set to by just calling return rather than return prefix.
	return
}

func main() {
	fmt.Println(Hello("Chris", ""))

}
