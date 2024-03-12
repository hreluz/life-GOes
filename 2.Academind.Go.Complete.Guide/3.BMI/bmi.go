package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	fmt.Println("Hello there")
	fmt.Println("-----------------------------")
	fmt.Print("Please enter your weight (kg): ")
	weightInput, _ := reader.ReadString('\n')

	fmt.Print("Please enter your height (m): ")
	heightInput, _ := reader.ReadString('\n') //this add a \n at the end

	fmt.Print(weightInput)
	fmt.Print(heightInput)

	// Save that user input in variables and remove the \n
	weightInput = strings.Replace(weightInput, "\n", "", -1)
	heightInput = strings.Replace(heightInput, "\n", "", -1)

	weight, _ := strconv.ParseFloat(weightInput, 64)
	height, _ := strconv.ParseFloat(heightInput, 32)

	fmt.Print(weight)
	fmt.Print(height)

}
