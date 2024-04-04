package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	//1
	hobbies := [5]string{"Play videogames", "watch movies", "cook", "read", "shopping"}
	fmt.Println(hobbies)

	//2
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	//3
	firstWay := hobbies[:2]
	secondWay := firstWay

	fmt.Println(firstWay)
	fmt.Println(secondWay)

	//4
	ff := firstWay[1:5]
	fmt.Println(ff)

	//5
	goals := []string{"finish this homework", "watch all the videos"}

	//6
	goals[1] = "watch all movies"
	goals = append(goals, "format laptop")
	fmt.Println(goals)

	//7
	products := []Product{Product{"product1", "1", 34.22}, Product{"product2", "2", 54.22}}
	products = append(products, Product{"product3", "3", 23.45})
	fmt.Println(products)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
