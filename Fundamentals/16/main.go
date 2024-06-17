package main

func sum(a, b *int) int {
	*a = 50
	*b = 30

	return *a + *b
}

func main() {

	myA := 10
	myB := 20

	println(myA, myB)
	total := sum(&myA, &myB)
	println(myA, myB, total)
}
