package main

import "fmt"

var userName = "Max"

func main() {
	shouldContinue := true

	if shouldContinue {
		userAge := 32

		fmt.Printf("Name: %v, Age: %v", userName, userAge)
	}

	printData()
}

func printData() {
	fmt.Println(userName)
}
