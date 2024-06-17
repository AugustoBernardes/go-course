package main

import "fmt"

var (
	b bool = true
	c int  = 3
	d string
	e float64
)

func main() {
	var myArray [3]int
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 4

	fmt.Println(len(myArray))
	fmt.Println(myArray[len(myArray)-1])

	for i, v := range myArray {
		fmt.Printf("O valor do indice %d é %d\n", i, v)
	}
}
