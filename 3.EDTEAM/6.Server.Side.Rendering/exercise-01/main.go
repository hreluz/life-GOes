package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl01 := "Hello World"
	t, err := template.New("exercise_1").Parse(tpl01)

	if err != nil {
		log.Fatalf("Error when parsing template %v", err)
	}

	err = t.Execute(os.Stdout, nil)

	if err != nil {
		log.Fatalf("Error when executing template %v", err)
	}
}
