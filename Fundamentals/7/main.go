package main

import "fmt"

func main() {
	salario := map[string]int{"Wesley": 1000, "Chico": 2000}
	println(salario["Chico"])
	delete(salario, "Chico")
	println(salario["Chico"])
	salario["Chico"] = 2222
	println(salario["Chico"])

	// Inicializar
	// salario2 := make(map[string]int)
	// salario2 := map[string]int{}

	for name, value := range salario {
		fmt.Printf("O Salario de %s é %d\n", name, value)
	}

	for _, value := range salario {
		fmt.Printf("O Salario de %d\n", value)
	}
}
