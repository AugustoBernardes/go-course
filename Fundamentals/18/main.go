package main

import "fmt"

func main() {
	var x interface{} = 10
	var y interface{} = "Hello World"

	printType(x)
	printType(y)
}

func printType(t interface{}) {
	fmt.Printf("O tipo %T , tem o valor %v\n", t, t)
}
