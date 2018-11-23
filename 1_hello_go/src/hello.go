package main

import "fmt"

const helloForDefault = "Stranger"
const langSpanish = "Spanish"
const langFrench = "French"
const helloEnglishPrefix = "Hello, "
const helloSpanishPrefix = "Hola, "
const helloFrenchPrefix = "Bonjour, "

// Maybe we need Hash(Lang -> GreetingPrefix)
func getPrefix(lang string) string {
	switch lang {
	case langSpanish :
		return helloSpanishPrefix
	case langFrench :
		return helloFrenchPrefix
	default:
		return helloEnglishPrefix
	}
}

func Hello(name string, lang string) string {
	prefix := getPrefix(lang)
	if name == "" {
		name = helloForDefault
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("World", "English"))
}
