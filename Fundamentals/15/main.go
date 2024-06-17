package main

func main() {
	a := 10

	var ponteiro *int = &a
	b := a

	println(*ponteiro)
	*ponteiro = 20
	println(a, *ponteiro, b)

}
