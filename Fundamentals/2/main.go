package main

import "fmt"

var (
	b bool = true
	c int  = 3
	d string
	e float64
)

func main() {
	fmt.Printf(`O tipo de e é %T`, b)
}
