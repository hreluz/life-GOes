package info

import "fmt"

const mainTitle = "BMI Calculator"
const separator = "-----------------------------"
const WeightPrompt = "Please enter your weight (kg): "
const HeightInput = "Please enter your height (m): "

func PrintWelcome() {
	fmt.Println(mainTitle)
	fmt.Println(separator)
}
