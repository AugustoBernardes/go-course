package main

import "fmt"

type Adress struct {
	street int
	city   string
}

// Assim está criado uma propriedade no objeto
type ClientType struct {
	name     string
	age      int
	sur_name string
	adrress  Adress
}

// Assim está composta
// type ClientType struct {
// 	name     string
// 	age      int
// 	sur_name string
// 	Adress
// }

func main() {

	client_data := ClientType{
		name:     "Augusto",
		age:      12,
		sur_name: "Bernardes",
	}

	client_data.adrress.city = "quirinopolis"

	fmt.Println(client_data)
}
