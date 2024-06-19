package main

import "fmt"

func main() {
	scores := make(map[string]int)
	scores["one"] = 1
	scores["two"] = 2
	scores["three"] = 3

	fmt.Println(scores)
}
