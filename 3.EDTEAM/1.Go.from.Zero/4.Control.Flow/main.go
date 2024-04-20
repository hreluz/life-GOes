package main

import "fmt"

func main() {

	fmt.Println("---------- IF, ELSE ------------")
	emoji := "😂"

	if emoji == "😂" {
		fmt.Println("it's a laughing emoticon")
	} else if emoji == "😭" {
		fmt.Println("it's a crying emoticon")
	} else {
		fmt.Println("the emoticon is: %s", emoji)
	}

	if emoji2 := "😂"; emoji2 == "😭" {
		fmt.Println("it's a crying emoticon")
	} else {
		fmt.Println("no emoticon found")
	}

	fmt.Println("\n\n---------- SWITCH ------------")
	emoji3 := "🐈"

	switch emoji3 {
	case "🐈":
		fmt.Println("It is a cat")
	case "🐶", "🐕":
		fmt.Println("It is a dog")
	default:
		fmt.Println("IT is not a domestic animal")
	}

	fmt.Println("\n\n---------- LOOPS ------------")

	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("\nSimilar to a while logic: ")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	fmt.Println("\nA forever loop: ")

	// for {
	// 	fmt.Println(i)
	// 	i++
	// }

	fmt.Println("\nFor range-slice: ")
	nums := []uint{2, 4, 6, 8}
	var v uint
	v = 1

	for i, value := range nums {
		fmt.Println("Index: ", i, "Value: ", value)
		nums[i] *= 2
		v = v * value
	}

	fmt.Println(v)
	fmt.Println(nums)

	fmt.Println("\nFor range-maps: ")
	sports := map[string]string{"basketball": "🏀", "football": "⚽"}
	for key, v := range sports {
		fmt.Println("Key", key, "Value", v)
	}

	fmt.Println("\nFor range-strings: ")
	hello := "hello"

	for _, v := range hello {
		fmt.Println(string(v))
	}
}
