package main

import "fmt"

type Adress struct {
	street int
	city   string
}

type ClientType struct {
	name      string
	age       int
	sur_name  string
	is_active bool
	adrress   Adress
}

type Person interface {
	disableClient()
}

func disable(person Person) {
	person.disableClient()
}

func (client ClientType) disableClient() {
	client.is_active = false
	fmt.Printf("Clinet is not active anymore")
	return
}

func main() {

	client_data := ClientType{
		name:      "Augusto",
		age:       12,
		sur_name:  "Bernardes",
		is_active: true,
	}

	client_data.adrress.city = "quirinopolis"

	disable(client_data)
}
