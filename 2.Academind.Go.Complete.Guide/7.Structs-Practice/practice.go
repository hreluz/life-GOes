package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

type Product struct {
	id          string
	title       string
	description string
	price       float64
}

func (p *Product) store() {
	file, _ := os.Create(p.id + ".txt")
	content := fmt.Sprintf(`
	Book id: %v
	Book title: %v
	Book description: %v
	Book price: %.2f`, p.id, p.title, p.description, p.price)
	file.WriteString(content)
}

var reader = bufio.NewReader(os.Stdin)

func main() {
	product1 := Product{
		"1",
		"Head Pattern Design",
		"Some designs well explained",
		23.44,
	}
	product2 := getProduct()
	product2.store()
	product1.print()
	product2.print()
}

func (product *Product) print() {
	fmt.Printf("Book id: %v\n", product.id)
	fmt.Printf("Book title: %v\n", product.title)
	fmt.Printf("Book description: %v\n", product.description)
	fmt.Printf("Book price: %.2f\n", product.price)
}

func getProduct() *Product {
	id := readUserInput("Please enter book id: ")
	name := readUserInput("Please enter book name: ")
	description := readUserInput("Please enter book description: ")
	priceInput := readUserInput("Please enter book price: ")
	price, _ := strconv.ParseFloat(priceInput, 64)
	return createProduct(id, name, description, price)
}

func readUserInput(promptText string) string {
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

func createProduct(id string, t string, d string, p float64) *Product {
	return &Product{id, t, d, p}
}

// Your Tasks
// 1) Create a new type / data structure for storing product data (e.g. about a book)
//		The data structure should contain these fields:
//		- ID
//		- Title / Name
//		- Short description
//		- price (number without currency, we'll just assume USD)
// 2) Create concrete instances of this data type in two different ways:
//		- Directly inside of the main function
//		- Inside of main, by using a "creation helper function"
//		(use any values for title etc. of your choice)
//		Output (print) the created data structures in the command line (in the main function)
// 3) Add a "connected function" that outputs the data + call that function from inside "main"
// 4) Change the program to fetch user input values for the different data fields
//		and create only one concrete instance with that entered data.
//		Output that instance data (via the connected function) then.
// 5) Bonus: Add a connected "store" function that writes that data into a file
//		The filename should be the unique id, the function should be called at the
//		end of main.
