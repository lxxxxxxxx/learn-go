package main

import (
	"fmt"
)

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello,"
const spanishHelloPrefix = "Hola,"
const frenchHelloPrefix = "Bonjour,"

//Hello :return Hello World.
func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func main() {
	fmt.Println(Hello("World", ""))
}

func greetingPrefix(language string)(prefix string){
	switch language{
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}