package mathTest

func Sum[T int | float64](a T, b T) T {
	return a + b
}

var A = 10

type Car struct {
	Type string
}

func (car Car) Start() string {
	return "Starting car"
}
