package main

import "fmt"

const englishHelloPrefix = "Let's begin "
const frenchHelloPrefix = "Commencer "

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}
	if lang == "" {
		lang = "english"
	}

	if lang == "french" {
		return frenchHelloPrefix + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Kirk", "french"))
}
