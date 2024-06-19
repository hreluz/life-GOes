package main

import (
	"errors"
	"fmt"
)

func division(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not possible at all")
	}

	return a / b, nil
}

func main() {
	res, err := division(3, 5)

	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println(res)
}
