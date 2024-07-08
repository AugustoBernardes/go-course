package main

import (
	"html/template"
	"os"
)

type Course struct {
	Name     string
	Duration int
}

func main() {
	curso := Course{"This course", 40}
	t := template.Must(template.New("Course template").Parse("Course:{{.Name}} , Duration: {{.Duration}}"))
	err := t.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
