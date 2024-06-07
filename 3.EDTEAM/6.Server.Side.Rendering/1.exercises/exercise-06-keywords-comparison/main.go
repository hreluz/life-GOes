package main

import (
	"log"
	"os"
	"text/template"
)

type numbers struct {
	X int
	Y int
}

func main() {
	tpl06 := `
Comparison
eq: {{ .X }} is equal to {{ .Y }}? = {{ eq .X .Y }}
ne: {{ .X }} is different to "!=" {{ .Y }}? = {{ ne .X .Y }}
lt: {{ .X }} is less than "<" {{ .Y }}? = {{ lt .X .Y }}
le: {{ .X }} is less or equal to "<=" {{ .Y }}? = {{ le .X .Y }}
gt: {{ .X }} is higher to ">" {{ .Y }}? = {{ gt .X .Y }}
ge: {{ .X }} is equal or equal to ">=" {{ .Y }}? = {{ ge .X .Y }}
`

	t, err := template.New("exercise_5").Parse(tpl06)

	if err != nil {
		log.Fatalf("Error when parsing template %v", err)
	}

	n := numbers{X: 5, Y: 6}

	err = t.Execute(os.Stdout, &n)

	if err != nil {
		log.Fatalf("Error when executing template %v", err)
	}
}
