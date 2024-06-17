package main

// Dessa forma estariamos duplicando coligos a toa
// func SumInter(m map[string]int) int {
// 	var sum int
// 	for _, value := range m {
// 		sum += value
// 	}

// 	return sum
// }

// func SumFloat(m map[string]float64) float64 {
// 	var sum float64
// 	for _, value := range m {
// 		sum += value
// 	}

// 	return sum
// }

type MyNumber int

type Number interface {
	~int | ~float64
}

func Sum[T Number](m map[string]T) T {
	var sum T
	for _, value := range m {
		sum += value
	}

	return sum
}

func Compare[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {

	mapInter := map[string]int{"Augusto": 100, "Luis": 10}
	mapFloat := map[string]float64{"Augusto": 100.10, "Luis": 10.20}
	mapInter2 := map[string]MyNumber{"Augusto": 100, "Luis": 10}

	// Antigo
	// println(SumInter(mapInter))
	// println(SumFloat(mapFloat))

	println(Sum(mapInter))
	println(Sum(mapInter2))
	println(Sum(mapFloat))
	println(Compare(10, 20))
}
