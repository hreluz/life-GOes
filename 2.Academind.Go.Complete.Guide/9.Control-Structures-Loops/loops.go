package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {

	selectedChoice, err := getUserChoice()

	if err != nil {
		fmt.Println("Invalid choice, exiting")
		return
	}

	switch selectedChoice {
	case "1":
		calculateSumUpToNumber()
	case "2":
		calculateFactorial()
	case "3":
		calculateSumManually()
	default:
		calculateListSum()
	}
}

func getInputNumber() (int, error) {
	inputNumber, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Invalid Input")
		return 0, err
	}

	inputNumber = strings.Replace(inputNumber, "\n", "", -1)
	chosenNumber, err := strconv.ParseInt(inputNumber, 0, 64)

	if err != nil {
		return 0, err
	}

	return int(chosenNumber), nil
}

func calculateSumUpToNumber() (int, error) {
	fmt.Print("Please enter your number: ")
	chosenNumber, err := getInputNumber()

	if err != nil {
		fmt.Println("Invalid number input")
		return 0, err
	}

	fmt.Println(chosenNumber)
	sum := 0
	for i := 1; i <= chosenNumber; i++ {
		sum = sum + i
	}

	fmt.Printf("Result: %v", sum)

	return chosenNumber, nil
}

func calculateFactorial() (int, error) {
	fmt.Print("Please enter your number: ")
	chosenNumber, err := getInputNumber()

	if err != nil {
		fmt.Println("Invalid number input")
		return 0, err
	}

	factorial := 1

	for i := 1; i <= chosenNumber; i++ {
		factorial *= i
	}

	fmt.Printf("Result: %v", factorial)
	return factorial, nil
}

func calculateListSum() {
	fmt.Println("Please enter a comma-separated list of numbers")

	inputNumberList, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Invalid number input")
		return
	}

	inputNumberList = strings.Replace(inputNumberList, "\n", "", -1)
	inputNumbers := strings.Split(inputNumberList, ",")

	sum := 0

	for index, value := range inputNumbers {
		fmt.Printf("Index: %v, Value: %v\n", index, value)
		number, err := strconv.ParseInt(value, 0, 64)

		if err != nil {
			continue
		}

		sum = sum + int(number)
	}
	fmt.Printf("Result: %v", sum)
}

func calculateSumManually() {
	isEnteringNumbers := true
	sum := 0
	fmt.Println("Keep entering numbers, the program will quit once you enter a character")

	for isEnteringNumbers {
		chosenNumber, err := getInputNumber()
		sum += chosenNumber
		isEnteringNumbers = err == nil
	}

	fmt.Printf("Result: %v", sum)
}

func getUserChoice() (string, error) {
	fmt.Println("Please make your choice")
	fmt.Println("1) Add up all the numbers of the number X")
	fmt.Println("2) Build the factorial up to number X")
	fmt.Println("3) Sum up manually entered numbers")
	fmt.Println("4) Sum up a list of entered numbers")

	fmt.Println("Please make your choice: ")
	userInput, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	userInput = strings.Replace(userInput, "\n", "", -1)

	if userInput == "1" || userInput == "2" || userInput == "3" || userInput == "4" {
		return userInput, nil
	} else {
		return "", errors.New("Invalid input!")
	}
}
