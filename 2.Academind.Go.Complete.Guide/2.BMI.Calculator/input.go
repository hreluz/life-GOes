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

func getUserMetrics() (float64, float64) {

	fmt.Print(info.WeightPrompt)
	weightInput, _ := reader.ReadString('\n')

	fmt.Print(info.HeightPrompt)
	heightInput, _ := reader.ReadString('\n')

	weightInput = strings.Replace(weightInput, "\n", "", -1)
	heightInput = strings.Replace(heightInput, "\n", "", -1)

	weight, _ := strconv.ParseFloat(weightInput, 64)
	height, _ := strconv.ParseFloat(heightInput, 64)

	return weight, height
}
