package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}
type hungarianDictionary struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	hd := hungarianDictionary{}

	printGreeting(eb)
	printGreeting(sb)
	printGreeting(hd)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// Custom logic for generating greeting
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	// Custom logic for generating spanish greeting
	return "Hola!"
}

func (hungarianDictionary) getGreeting() string {
	return "Szia!"
}
