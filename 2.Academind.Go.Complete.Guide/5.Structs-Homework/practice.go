package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	id          string
	title       string
	description string
	price       float64
}

func (product *Product) store() {
	file, _ := os.Create(product.id + ".txt")

	content := fmt.Sprintf(
		"ID: %v\nTitle: %v\nDescription: %v\nPrice: %.2f",
		product.id,
		product.title,
		product.description,
		product.price,
	)
	file.WriteString(content)

	file.Close()
}

func (product *Product) printData() {
	fmt.Printf("ID: %v\n", product.id)
	fmt.Printf("Title: %v\n", product.title)
	fmt.Printf("Description: %v\n", product.description)
	fmt.Printf("Price: %.2f\n", product.price)
}

func createProduct(id string, title string, description string, price float64) *Product {
	return &Product{
		id,
		title,
		description,
		price,
	}
}

var reader = bufio.NewReader(os.Stdin)

func main() {

	firstProduct := Product{
		"first-product",
		"a book",
		"a nice book",
		10.99,
	}

	secondProduct := *createProduct("second-product", "A notebook", "a nice notebook", 41.2)
	thirdProduct := getProduct()
	firstProduct.printData()
	secondProduct.printData()
	thirdProduct.printData()

	thirdProduct.store()
}

func readUserInput(reader *bufio.Reader, prompText string) string {
	fmt.Print(prompText)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)

	return userInput
}

func getProduct() *Product {
	fmt.Println("please enter the product data.")
	fmt.Println("-------------------------------------")
	reader = bufio.NewReader(os.Stdin)

	idInput := readUserInput(reader, "Product ID: ")
	titleInput := readUserInput(reader, "Product title: ")
	descriptionInput := readUserInput(reader, "Product Description: ")
	priceInput, _ := strconv.ParseFloat(readUserInput(reader, "Product Price: "), 64)

	return &Product{
		idInput,
		titleInput,
		descriptionInput,
		priceInput,
	}
}

// Your Tasks

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
