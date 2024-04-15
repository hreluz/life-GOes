package main

import (
	"bufio"
	"fmt"
	"os"
)

type Storer interface {
	Store(fileName string)
}

type Prompter interface {
	PromptForInput()
}

type PrompterStore interface {
	Prompter
	Storer
}

type userInputData struct {
	input string
}

func newUserInputData() *userInputData {
	return &userInputData{}
}

func (usr *userInputData) PromptForInput() {
	fmt.Print("Your input please: ")

	reader := bufio.NewReader(os.Stdin)

	userInput, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Fetching user input failed!")
		return
	}

	usr.input = userInput
}

func (usr *userInputData) Store(fileName string) {
	file, err := os.Create(fileName)

	if err != nil {
		fmt.Println("Creating the file failed!")
		return
	}

	defer file.Close()

	file.WriteString(usr.input)
}

func main() {
	data := newUserInputData()
	data.PromptForInput()
	data.Store("user1.txt")

	handleUserInput(data)
}

func handleUserInput(container PrompterStore) {
	fmt.Println("Ready to store your data")
	container.PromptForInput()
	container.Store("container.txt")
}
