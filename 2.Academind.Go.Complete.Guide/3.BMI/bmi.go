package main

import (
	"fmt"
	// found in mod when creating with go mod init
	"github.com/hreluz/bmi/info"
)

func main() {
	info.PrintWelcome()

	weight, height := getUserMetrics()

	// Save that user input in variables and remove the \n

	// fmt.Print(weight)
	// fmt.Print(height)

	// Calculate the BMI (weight/ (height*height))
	bmi := weight / (height * height)

	// Output the calculated BMI
	fmt.Printf("Your BMI is : %.2f", bmi)
}
