package main

import "fmt"

func main() {

	var myVar interface{} = "Augusto"
	println(myVar.(string))

	// Usando ok é possivel validar se conversão é possivel
	res, ok := myVar.(int)
	fmt.Printf("A resposta de %v é %v", res, ok)

	// Aqui é pra custoso, e gera Panic pois ação não pode ser feita
	res2 := myVar.(int)
	println(res2)
}
