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
	tmp := template.New("Course templatetmp")
	tmp, _ = tmp.Parse("Course:{{.Name}} , Duration: {{.Duration}}")
	err := tmp.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
