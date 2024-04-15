package main

import (
	"fmt"
	"io/fs"
	"os"
)

type logger interface {
	log()
	// pritMessage(message string) int
}

type logData struct {
	message  string
	fileName string
}

func (lData *logData) log() {
	err := os.WriteFile(lData.fileName, []byte(lData.message), fs.FileMode(0644))

	if err != nil {
		fmt.Println("Storing the log data failed!")
		message := err.Error()
		fmt.Println(message)
	}
}

type loggableString string

func (text loggableString) log() {
	fmt.Println(text)
}

func main() {
	userLog := logData{"User logged in", "user-log.txt"}
	// do more work
	// userLog.log()
	createLog(&userLog)

	message := loggableString("[INFO] User created!")
	// do more work
	createLog(message)
	// createLog("Test")

	outputValue(message)
	outputValue(userLog)

	fmt.Println(double(2))
	fmt.Println(double(2.5))

	numbers := []int{1, 2, 3}
	fmt.Println(double(numbers))
}

func createLog(data logger) {
	// more things we do
	data.log()
}

func outputValue(value interface{}) {
	// assert type of a value
	val, ok := value.(loggableString)
	fmt.Println(val, ok)
	fmt.Println(value)
}

func double(value interface{}) interface{} {

	switch val := value.(type) {
	case int:
		return val * 2
	case float64:
		return val * 2
	case []int:
		newNumbers := append(val, val...)
		return newNumbers
	default:
		return "case not covered"
	}
}
