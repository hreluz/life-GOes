package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	var productNames [4]string
	productNames = [4]string{"A book"}

	prices := [4]float64{4.3, 2.3, 1.0, 99.3}
	fmt.Println(prices)
	fmt.Println(productNames)
	fmt.Println(prices[1])
}
