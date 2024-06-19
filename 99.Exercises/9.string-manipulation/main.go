package main

import "fmt"

func reverseWord(s string) string {
	i := len(s) - 1
	r := ""

	for i >= 0 {
		r += string(s[i])
		i--
	}

	return r
}

func main() {
	s := "batman"
	fmt.Printf("The original string is: %s\n", s)
	fmt.Printf("The reverse word is: %s", reverseWord(s))
}
