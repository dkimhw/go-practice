package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	if lang == "English" {
		return englishHelloPrefix + name
	}

	if lang == "Spanish" {
		return spanishHelloPrefix + name
	}

	return frenchHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Chris", "English"))
}
