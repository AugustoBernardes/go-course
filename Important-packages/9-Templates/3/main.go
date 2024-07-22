package main

import (
	"os"
	"text/template"
)

type Course struct {
	Name     string
	Duration int
}

type Courses []Course

func main() {
	t := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := t.Execute(os.Stdout, Courses{
		{"go", 12},
		{"js", 13},
		{"pyhon", 11},
	})
	if err != nil {
		panic(err)
	}
}
