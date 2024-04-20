package main

import "fmt"

func main() {
	fmt.Println("------------- SLICES -------------")
	// slices (slices do not have data, they are just pointers to arrays)
	set := [7]string{"🦁", "🐴", "🐮", "🦋", "🐦", "🛩️", "🚁"}

	animals := set[0:5]

	fmt.Println("Array:", set)
	fmt.Println("Animals slice:", animals)
	fly := set[3:7]
	fmt.Println("Fly slice:", fly)
	fly[0] = "🦅"

	fmt.Println("\nArray:", set)
	fmt.Println("Animals slice:", animals)
	fmt.Println("Fly slice:", fly)

	food := [5]string{"🌭", "🍓", "🍋", "🍔", "🍕"}
	fruits := food[1:3]

	fmt.Println("Array food:", food)
	fmt.Println("Slice fruits:", fruits)
	fmt.Println("Slice fruits length:", len(fruits))
	// capacity will get the size from the index start until the end   {"🍓", "🍋", "🍔", "🍕"}
	fmt.Println("Slice fruits capacity:", cap(fruits))

	fruits = append(fruits, "🍍", "🍉")
	fmt.Println("\nArray food:", food)
	fmt.Println("Same slice fruits:", fruits)
	fmt.Println("Slice fruits capacity:", cap(fruits))

	fruits = append(fruits, "🍐")
	fmt.Println("\nArray food:", food)
	fmt.Println("Same slice fruits:", fruits)
	fmt.Println("Slice fruits capacity:", cap(fruits))

	// declaring a slice
	fruits2 := []string{"🍇", "🍑"}
	fruits2 = append(fruits2, "🍌")
	fmt.Println("\nfruits", fruits2)
	fmt.Println("size", len(fruits2))
	fmt.Println("capacity", cap(fruits2))

	fruits3 := make([]string, 0, 3)
	fmt.Println("\nfruits:", fruits3)
	fmt.Println("size", len(fruits3))
	fmt.Println("capacity", cap(fruits3))
	fruits3 = append(fruits, "🍓", "🍋")
	fmt.Println("fruits:", fruits3)
	fmt.Println("size", len(fruits3))
	fmt.Println("capacity", cap(fruits3))
}
