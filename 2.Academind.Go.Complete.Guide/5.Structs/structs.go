package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

type User struct {
	firstName   string
	lastName    string
	birthdate   string
	createdDate time.Time
}

func (user *User) outputDetails() {
	fmt.Printf("My name is %v %v (born on %v)", user.firstName, user.lastName, user.birthdate)
}

func NewUser(fName string, lName string, bDate string) *User {
	created := time.Now()

	user := User{
		firstName:   fName,
		lastName:    lName,
		birthdate:   bDate,
		createdDate: created,
	}

	return &user
}

var reader = bufio.NewReader(os.Stdin)

func main() {
	var newUser *User

	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")
	// created := time.Now()

	// newUser = User{
	// 	firstName:   firstName,
	// 	lastName:    lastName,
	// 	birthdate:   birthdate,
	// 	createdDate: created,
	// }

	// newUser = User{}
	// newUser = User{firstName, lastName, birthdate, created}

	newUser = NewUser(firstName, lastName, birthdate)

	fmt.Println(newUser)

	// outputUserDetails(newUser)
	newUser.outputDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')

	var cleanedInput string

	if runtime.GOOS == "windows" {
		cleanedInput = strings.Replace(userInput, "\r\n", "", -1)
	} else {
		cleanedInput = strings.Replace(userInput, "\n", "", -1)
	}
	return cleanedInput

}
