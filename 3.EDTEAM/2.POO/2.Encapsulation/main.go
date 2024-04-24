package main

import (
	"fmt"

	"github.com/hreluz/encapsulation/course"
)

func main() {
	Go := course.New("Go From Zero", 0, false)
	Go.SetUserIDs([]uint{12, 56, 89})
	Go.SetClasses(
		map[uint]string{
			1: "Introduction",
			2: "Structs",
			3: "Maps",
		},
	)

	Go.SetName("POO with GO")
	fmt.Println(Go.Price())
	Go.PrintClasses()
	fmt.Printf("%+v", Go)
	fmt.Println(Go.Name())
}
