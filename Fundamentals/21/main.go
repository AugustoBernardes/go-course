package main

import (
	"fmt"
	mathTest "package-class/math"

	"github.com/google/uuid"
)

func main() {
	sum := mathTest.Sum(10, 20)

	car := mathTest.Car{Type: "fiat"}
	fmt.Println(car.Start())
	fmt.Println("Sum resul:", sum)
	fmt.Println(mathTest.A)
	fmt.Println(uuid.New())
}
