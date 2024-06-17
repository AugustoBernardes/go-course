package main

const a string = "Hello World"

var (
	b bool = true
	c int  = 3
	d string
	e float64
)

func main() {
	d = "teste"
	// Outra forma de declarar variavel
	f := "X"
	println(a, b, c, d, e, f)
}
