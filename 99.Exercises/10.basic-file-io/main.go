package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	content := "this is the content"
	filename := "somefile.txt"
	err := ioutil.WriteFile(filename, []byte(content), 0644)

	if err != nil {
		fmt.Println("An error occured when writing on file: ", err)
	}

	readContent, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("An error occured when reading the file: ", err)
	}

	fmt.Println("File content", string(readContent))
}
