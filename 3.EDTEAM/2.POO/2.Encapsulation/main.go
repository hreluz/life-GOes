package main

import (
	"fmt"

	"github.com/hreluz/encapsulation/course"
)

func main() {
	Go := course.New("Go From Zero", 0, false)

	Go.UserIDs = []uint{12, 56, 89}
	Go.Classes = map[uint]string{
		1: "Introduction",
		2: "Structs",
		3: "Maps",
	}

	fmt.Println(Go.Price)
	Go.PrintClasses()
	fmt.Printf("%+v", Go)
}
