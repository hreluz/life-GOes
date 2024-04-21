package main

import (
	"fmt"

	"github.com/hreluz/packages-modules/greet"

	"rsc.io/quote"

	// using different version of the same package
	quoteV3 "rsc.io/quote/v3"
)

func main() {
	fmt.Println("Hello")

	englishGreeting := greet.English()

	italianGreeting := greet.Italian()

	spanishGreeting := greet.Spanish()

	fmt.Println(englishGreeting)
	fmt.Println(italianGreeting)
	fmt.Println(spanishGreeting)
	fmt.Println(quote.Hello())
	fmt.Println(quoteV3.Concurrency())
}
