package main

import "fmt"

func main() {

	var example = 3
	fmt.Println(example)

	var example2 int = 4
	fmt.Println(example2)

	var example3 int
	example3 = 5

	fmt.Println(example3)

	var a, b, c = 1, 2, "example"

	fmt.Println(a, b, c)

	var (
		d int    = 1
		e int    = 3
		f string = "hello"
	)

	fmt.Println(d, e, f)

	ca := 3
	fmt.Println(ca)

	var myName string = "Bruce"
	fmt.Println("My name is", myName)

	var lastName = "Wayne"
	fmt.Println("My last name is", lastName)

	username := "batman"
	fmt.Println("username =", username)

	part1, other := 1, 5
	fmt.Println("part 1 is", part1, "other is", other)

	part2, other := 2, 0
	fmt.Println("part 2 is", part2, "other is", other)

	var (
		lessonName = "Variables"
		lessonType = "Demo"
	)

	fmt.Println("lessonName=", lessonName)
	fmt.Println("lessonType=", lessonType)

}
