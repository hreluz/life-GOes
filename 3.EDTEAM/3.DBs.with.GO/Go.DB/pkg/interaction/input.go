package interaction

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/hreluz/go-db/storage"
)

var reader = bufio.NewReader(os.Stdin)

func GetDBChoice() storage.Driver {

	for {
		playerChoice, _ := GetUserInput("Your db choice: ")
		switch playerChoice {
		case "1":
			return storage.MySQL
		case "2":
			return storage.Postgres
		default:
			os.Exit(1)
		}
	}
}

func GetActionChoice() storage.MENU_OPTIONS {

	for {
		playerChoice, _ := GetUserInput("Your action choice: ")

		switch playerChoice {
		case "1":
			return storage.CREATE_PRODUCT
		case "2":
			return storage.UPDATE_PRODUCT
		case "3":
			return storage.DELETE_PRODUCT
		case "4":
			return storage.SHOW_ALL_PRODUCT
		case "5":
			return storage.SHOW_BY_ID_PRODUCT
		case "6":
			return storage.CREATE_INVOICE
		case "7":
			return storage.SHOW_INVOICES
		case "8":
			return storage.MIGRATE_ALL
		default:
			os.Exit(1)
		}

		fmt.Println("Fetching the user input failed, Please try again")
	}
}

func GetUserInput(prompt string) (string, error) {
	if prompt != "" {
		fmt.Printf("%s: ", prompt)
	}

	userInput, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	userInput = strings.Replace(userInput, "\n", "", -1)
	return userInput, nil
}

func GetInputNumber() uint {
	inputNumber, err := reader.ReadString('\n')

	if err != nil {
		return 0
	}

	inputNumber = strings.Replace(inputNumber, "\n", "", -1)
	chosenNumber, err := strconv.ParseInt(inputNumber, 0, 64)

	if err != nil {
		return 0
	}

	return uint(chosenNumber)
}
