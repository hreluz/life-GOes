package lists

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	prices := []float64{10.99, 8.88}
	fmt.Println(prices[0:1])

	prices[1] = 44.2
	fmt.Println(prices)

	updatedPrices := append(prices, 5.99)
	fmt.Println(updatedPrices, prices)

	prices = append(prices, 12.3)
	fmt.Println(prices)
	prices = prices[1:]
	fmt.Println(prices)

	discountPrices := []float64{102.99, 21.24, 43.33}
	prices = append(prices, discountPrices...)
	fmt.Println(prices)
}

// func main() {
// 	var productNames [4]string
// 	productNames = [4]string{"A book"}

// 	prices := [4]float64{4.3, 2.3, 1.0, 99.3}
// 	fmt.Println(prices)

// 	productNames[2] = "A carpet"

// 	fmt.Println(productNames)
// 	fmt.Println(prices[1])

// 	// first value is
// 	featurePrices := prices[1:3]
// 	featurePrices[0] = 199.99
// 	fmt.Println(featurePrices)

// 	featurePrices2 := prices[:4]
// 	fmt.Println(featurePrices2)

// 	highlightedPrices := featurePrices2[:2]
// 	fmt.Println(highlightedPrices)

// 	fmt.Println(prices)
// 	fmt.Println(len(featurePrices), cap(featurePrices))

// 	fmt.Println(highlightedPrices, len(highlightedPrices), cap(highlightedPrices))
// 	highlightedPrices = highlightedPrices[:3]
// 	fmt.Println(highlightedPrices, len(highlightedPrices), cap(highlightedPrices))
// }
