package main

import (
	"html/template"
	"log"
	"os"
)

type person struct {
	Name     string
	Age      int
	password string
}

func (p *person) Greeting() string {
	return "Hello, I am " + p.Name
}

func (p *person) Add2(a int) int {
	return a + 2
}

func (p *person) Multiply(a, b int) int {
	return a * b
}
func main() {
	tpl02 := `Hello {{ .Name }}, his greeting es: {{ .Greeting }}
	Sum 2 to {{ .Age }}, the result is {{ .Add2 .Age }}
	Multiply  {{ .Age }}, * 6, the result is {{ .Multiply .Age 6 }}
`
	t, err := template.New("exercise_3").Parse(tpl02)

	if err != nil {
		log.Fatalf("Error when parsing template %v", err)
	}

	p := person{"Batman", 35, "123456"}
	err = t.Execute(os.Stdout, &p)

	if err != nil {
		log.Fatalf("Error when executing template %v", err)
	}
}
