package main

import "fmt"

func main() {
	Go := Course{
		Name:    "Go from zero",
		Price:   12.35,
		IsFree:  false,
		UserIDs: []uint{12, 56, 89},
		Classes: map[uint]string{
			1: "Introduction",
			2: "Structs",
			3: "Maps",
		},
	}

	fmt.Println(Go.Price)
	Go.PrintClasses()
	Go.ChangePrice(65.22) // -> (&Go).ChangePrice(65.22)
	fmt.Println(Go.Price)
	// fmt.Println(Go.Name)

	PHP := Course{
		"PHP from zero",
		11.35,
		false,
		[]uint{12, 56, 89},
		map[uint]string{
			1: "Introduction",
			2: "Classes",
			3: "Interfaces",
		},
	}

	PHP.PrintClasses()
	// fmt.Println(PHP.Name)

	CSS := Course{Name: "CSS from Zero"}
	fmt.Println(CSS)

	JS := Course{}
	JS.Name = "Javascript Course"
	JS.UserIDs = []uint{12, 67}

	fmt.Println(JS)
}
