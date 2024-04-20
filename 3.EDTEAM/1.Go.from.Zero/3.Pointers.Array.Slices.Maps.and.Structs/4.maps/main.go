package main

import "fmt"

func main() {
	fmt.Println("------------- MAPS -------------")
	animals := make(map[string]string)
	animals["cat"] = "😺"
	animals["dog"] = "🐶"
	fmt.Println(animals)

	fruits := map[string]string{
		"apple":  "🍎",
		"banana": "🍌",
	}

	fmt.Println(fruits)
	delete(fruits, "banana")
	fmt.Println(fruits)

	fmt.Println(animals["cat"])
	fmt.Println(animals["gorilla"])

	if _, ok := animals["gorilla"]; !ok {
		animals["gorilla"] = "🦍"
	}

	fmt.Println(animals)
}
