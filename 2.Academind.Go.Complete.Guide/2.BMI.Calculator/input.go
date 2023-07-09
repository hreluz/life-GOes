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
	weight := getUserInfo(info.WeightPrompt)
	height := getUserInfo(info.HeightPrompt)

	return weight, height
}

func getUserInfo(prompText string) (value float64) {

	fmt.Print(prompText)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)
	value, _ = strconv.ParseFloat(userInput, 64)

	return
}
