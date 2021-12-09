package main

import "fmt"

func getPrefix(language string) string {
	prefixMap := map[string]string {
		"english": "Hello ",
		"spanish": "Hola ",
		"french": "Bonjour ",
	}

	if language == "" {
		return prefixMap["english"]
	}

	return prefixMap[language]
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return getPrefix(language) + name
}

func main() {
	fmt.Println(Hello("World", "english"))
}