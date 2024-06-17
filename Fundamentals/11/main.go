package main

import "fmt"

type ClientType struct {
	name     string
	age      int
	sur_name string
}

func main() {

	client_data := ClientType{
		name:     "Augusto",
		age:      12,
		sur_name: "Bernardes",
	}

	fmt.Println(client_data)
}
