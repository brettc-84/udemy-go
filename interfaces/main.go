package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	// sb := spanishBot{}

	printGreeting(eb)
	// printGreeting(sb)
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGretting())
}

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGretting())
// }

// when not using the receiver variable, it can be ommitted
func (englishBot) getGretting() string {
	// custom logic for english bot implementation
	return "Hello there"
}

func (spanishBot) getGretting() string {
	// custom logic for spanish bot implementation
	return "Ola!"
}
