package main

import "fmt"

func main() {


	// Slice
	numbers := []string{"one", "two", "three"}
	for key, value := range numbers {
		fmt.Println(key, value)
	}

	// Almost same as while
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// Infinite loop
	for {
		fmt.Println("Hello world")
	}
}
