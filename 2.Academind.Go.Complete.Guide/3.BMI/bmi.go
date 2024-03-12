package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

const mainTitle = "BMI Calculator"
const separator = "-----------------------------"
const weightPrompt = "Please enter your weight (kg): "
const heightInput = "Please enter your height (m): "

func main() {
	fmt.Println(mainTitle)
	fmt.Println(separator)
	fmt.Print(weightPrompt)
	weightInput, _ := reader.ReadString('\n')

	fmt.Print(heightInput)
	heightInput, _ := reader.ReadString('\n') //this add a \n at the end

	fmt.Print(weightInput)
	fmt.Print(heightInput)

	// Save that user input in variables and remove the \n
	weightInput = strings.Replace(weightInput, "\n", "", -1)
	heightInput = strings.Replace(heightInput, "\n", "", -1)

	weight, _ := strconv.ParseFloat(weightInput, 64)
	height, _ := strconv.ParseFloat(heightInput, 32)

	// fmt.Print(weight)
	// fmt.Print(height)

	// Calculate the BMI (weight/ (height*height))
	bmi := weight / (height * height)

	// Output the calculated BMI
	fmt.Printf("Your BMI is : %.2f", bmi)
}
