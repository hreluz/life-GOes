package main

import (
	"fmt"
)

func main() {
	data := 1

	//child
	go func() {
		data++
	}()

	fmt.Println(data)
}
