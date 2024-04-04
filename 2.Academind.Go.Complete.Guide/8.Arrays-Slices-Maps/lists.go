package main

import "fmt"

func main() {
	// var productNames [4]string
	var productNames [4]string = [4]string{"A book", "two books"}
	productNames[2] = "A carpet"

	// arrays
	prices := [4]float64{10.99, 9.33, 405.22, 23}

	fmt.Println(prices[3])
	fmt.Println(productNames)

	// slices reference to initial array
	featuredPrices := prices[1:3]
	highlightedPrices := featuredPrices[:1]
	prices[1] = 34444
	fmt.Println(featuredPrices)
	fmt.Println(highlightedPrices)

	// cap shows the capacity of the field
	// capacity counts from the start element, but cannot go before that
	fmt.Println(len(featuredPrices), cap(featuredPrices), cap(highlightedPrices))

	highlightedPrices = highlightedPrices[:3]
	fmt.Println(highlightedPrices)
	fmt.Println(len(highlightedPrices), cap(featuredPrices), cap(highlightedPrices))
}
