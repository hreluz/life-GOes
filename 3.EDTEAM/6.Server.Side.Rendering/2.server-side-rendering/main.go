package main

import (
	"fmt"
	"html/template"
)

func main() {
	myFunctions := template.FuncMap{
		"usd": func(a float64) string {
			return fmt.Sprintf("$%.fUSD", a)
		},
	}
}
