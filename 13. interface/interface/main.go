package main

import "fmt"

type Bot interface {
	getGreeting() string
}

type EnglishBot struct{}

type SpanishBot struct{}

func main() {
	eb := EnglishBot{}
	sb := SpanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b Bot) {
	fmt.Println(b.getGreeting())
}

func (EnglishBot) getGreeting() string {
	return "Hi There!"
}

func (SpanishBot) getGreeting() string {
	return "Hola!"
}
