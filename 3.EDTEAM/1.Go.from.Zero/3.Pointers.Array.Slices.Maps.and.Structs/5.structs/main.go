package main

import "fmt"

type course struct {
	Name      string
	Professor string
	Country   string
}

func main() {
	fmt.Println("------------- STRUCTS -------------")

	db := course{
		Name:      "Database",
		Professor: "Linus Torvalds",
		Country:   "USA",
	}

	git := course{"Git", "Someone", "Peru"}

	fmt.Printf("%+v\n", db)
	fmt.Printf("%+v\n", git)
	fmt.Printf("%+v\n", db.Name)
	fmt.Printf("%+v\n", git.Name)

	p := &db
	fmt.Printf("%+v\n", db)
	fmt.Printf("%+v\n", p)
	p.Professor = "Einstein"
	fmt.Printf("%+v\n", db)
	fmt.Printf("%+v\n", p)
}
