package main

import (
	"fmt"
	"strconv"
	"strings"
	// found in mod when creating with go mod init
	"github.com/hreluz/bmi/info"
)

func main() {
	info.PrintWelcome()
	fmt.Print(info.WeightPrompt)
	weightInput, _ := reader.ReadString('\n')

	fmt.Print(info.HeightInput)
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
