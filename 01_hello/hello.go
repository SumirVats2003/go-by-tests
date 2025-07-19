package main

import "fmt"

const (
  spanish = "Spanish"
  french = "French"

  englishHelloPrefix = "Hello"
  spanishHelloPrefix = "Hola"
  frenchHelloPrefix = "Bonjour"
)

func Hello(name, lang string) string {
  if name == "" {
    name = "World"
  }
  prefix := getPrefix(lang)

  return prefix + " " + name
}

func getPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
  default:
    prefix = englishHelloPrefix
	}
	return
}

func main() {
  fmt.Println(Hello("world", ""))
}
