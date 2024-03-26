package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/hreluz/bmi/info"
)

var reader = bufio.NewReader(os.Stdin)

func getUserMetrics() (weight float64, height float64) {

	fmt.Print(info.WeightPrompt)
	weightInput, _ := reader.ReadString('\n')

	fmt.Print(info.HeightInput)
	heightInput, _ := reader.ReadString('\n') //this add a \n at the end

	fmt.Print(weightInput)
	fmt.Print(heightInput)

	weightInput = strings.Replace(weightInput, "\n", "", -1)
	heightInput = strings.Replace(heightInput, "\n", "", -1)

	weight, _ = strconv.ParseFloat(weightInput, 64)
	height, _ = strconv.ParseFloat(heightInput, 32)

	return
}
