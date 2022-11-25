package main

import "fmt"

func main() {
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf756",
	// }
	// colors := map[string]string{}
	// colors := make(map[string]string)
	colors := make(map[int]string)

	colors[10] = "#ffff"
	colors[20] = "#00000"

	delete(colors, 10)

	fmt.Println(colors)
}
