package main

import "fmt"

func main() {
	// slices array
	prices := []float64{10, 33}
	fmt.Println(prices[0:1])

	prices[1] = 9.99

	updatedPrices := append(prices, 5.34)
	prices = prices[1:]
	fmt.Println(updatedPrices)
	fmt.Println(prices)

	prices2 := []float64{10.99, 9.33, 405.22, 23}
	discountPrices := []float64{2, 3, 5, 6}
	prices2 = append(prices2, discountPrices...)
	fmt.Println(prices2)

}

// func main() {
// fixed length array
// 	// var productNames [4]string
// 	var productNames [4]string = [4]string{"A book", "two books"}
// 	productNames[2] = "A carpet"

// 	// arrays
// 	prices := [4]float64{10.99, 9.33, 405.22, 23}

// 	fmt.Println(prices[3])
// 	fmt.Println(productNames)

// 	// slices reference to initial array
// 	featuredPrices := prices[1:3]
// 	highlightedPrices := featuredPrices[:1]
// 	prices[1] = 34444
// 	fmt.Println(featuredPrices)
// 	fmt.Println(highlightedPrices)

// 	// cap shows the capacity of the field
// 	// capacity counts from the start element, but cannot go before that
// 	fmt.Println(len(featuredPrices), cap(featuredPrices), cap(highlightedPrices))

// 	highlightedPrices = highlightedPrices[:3]
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(len(highlightedPrices), cap(featuredPrices), cap(highlightedPrices))
// }
