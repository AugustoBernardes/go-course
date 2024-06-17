package main

import (
	"errors"
	"fmt"
)

func main() {
	value, error := sum(50, 2)

	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(value)
}

func sum(a, b int) (int, error) {
	sumTotal := a + b

	if sumTotal > 50 {
		return sumTotal, errors.New("é maior que 50")
	}
	return sumTotal, nil
}
