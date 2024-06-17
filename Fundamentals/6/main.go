package main

import "fmt"

func main() {
	s := []int{10, 20, 30, 50, 60}

	fmt.Printf("len=%d cap=%d %v \n", len(s), cap(s), s)

	fmt.Printf("len=%d cap=%d %v \n", len(s[:0]), cap(s[:0]), s[:0])

	//  Antes dos :

	fmt.Printf("len=%d cap=%d %v \n", len(s[:3]), cap(s[:3]), s[:3])

	fmt.Printf("len=%d cap=%d %v \n", len(s[:5]), cap(s[:5]), s[:5])

	//  Depois dos :
	//  Quando slice é passado apenas a 1 propriedade antes dos : a capacidade é alterada

	fmt.Printf("len=%d cap=%d %v \n", len(s[3:]), cap(s[3:]), s[3:])

	fmt.Printf("len=%d cap=%d %v \n", len(s[5:]), cap(s[5:]), s[5:])

	fmt.Printf("len=%d cap=%d %v \n", len(s[0:]), cap(s[0:]), s[0:])

	s = append(s, 110)

	fmt.Printf("len=%d cap=%d %v \n", len(s), cap(s), s)
}
