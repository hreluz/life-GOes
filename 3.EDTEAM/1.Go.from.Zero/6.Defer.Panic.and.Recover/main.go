package main

import (
	"fmt"
	"os"
)

func main() {
	// deferIntro()

	// deferExampleRealCase()

	panicFn()

}

func deferIntro() {
	// it executes it, at the end of the function
	defer fmt.Println(3)

	fmt.Println(1)
	fmt.Println(2)
	defer fmt.Println(4)
	defer fmt.Println(5)

	a := 5
	// this will execute it as 5, because at this time, it has this value, it does not check for the value after
	defer fmt.Println("Defer a:", a)

	a = 10
	fmt.Println("Before defer a: ", a)
}

func deferExampleRealCase() {
	file, err := os.Create("hello.txt")

	if err != nil {
		fmt.Printf("There was an error when creating the file %v", err)
		return
	}

	defer file.Close()

	_, err = file.Write([]byte("Hello there!"))

	if err != nil {
		// file.Close()
		fmt.Printf("There was an error when writing the file %v", err)
		return
	}

	// file.Close()
}

func panicFn() {
	division(10, 2)
	division(20, 5)
	division(5, 0)
	division(10, 4)
}

func division(dividen, divisor int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovering from panic ", r)
		}
	}()

	validateDivisor(divisor)
	fmt.Println(dividen / divisor)
}

func validateDivisor(divisor int) {
	if divisor == 0 {
		panic("Can't do a division by zero")
	}
}
