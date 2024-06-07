package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

func main() {
	t, err := template.New("").ParseFiles("./invitation.tpl", "./confirmation.tpl")

	if err != nil {
		log.Fatalf("Error when parsing template %v", err)
	}

	p := person{"Batman", 121}

	err = t.ExecuteTemplate(os.Stdout, "invitation.tpl", &p)

	if err != nil {
		log.Fatalf("Error when executing template %v", err)
	}

	if p.Age >= 18 {
		err = t.ExecuteTemplate(os.Stdout, "confirmation.tpl", &p)

		if err != nil {
			log.Fatalf("Error when executing template %v", err)
		}
	}
}
