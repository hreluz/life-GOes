package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"text/template"
)

func main() {
	myFunctions := template.FuncMap{
		"addOne": func(a int) int {
			a++
			return a
		},
		"lower": func(a string) string {
			return strings.ToLower(a)
		},
		"hours": func(a, b int) string {
			return "the hours are: " + strconv.Itoa((a+1)*b)
		},
	}

	superheroes := []string{
		"Batman",
		"Superman",
		"Flash",
	}

	tpl09 := `
Best Superheroes
{{- range $i, $course := .}}
{{ addOne $i }} {{ lower $course }}, {{ hours $i 3}}
{{- else }}
There are no elements
{{- end }}
`

	t, err := template.New("exercise_9").Funcs(myFunctions).Parse(tpl09)

	if err != nil {
		log.Fatalf("Error when parsing template %v", err)
	}

	err = t.Execute(os.Stdout, superheroes)

	if err != nil {
		log.Fatalf("Error when executing template %v", err)
	}
}
