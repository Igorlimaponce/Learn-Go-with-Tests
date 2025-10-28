package main

import "fmt"

func main() {

	fmt.Printf("%s", Hello("Joao", "Spanish"))
}

const HelloPrefix = "Hello,"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(name, language)
}

func greetingPrefix(name string, language string) string {
	switch language {
	case "Spanish":
		return "Hola, " + name
	case "French":
		return "Bonjour, " + name
	case "":
		return fmt.Sprintf("%s %s", HelloPrefix, name)
	default:
		return fmt.Sprintf("%s %s %s", HelloPrefix, name, language)
	}
}
