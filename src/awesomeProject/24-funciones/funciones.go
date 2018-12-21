package main

import (
	"fmt"
)

func imprimirNombre(name string) {
	fmt.Println("Fuera del main")
	fmt.Println("El name es: ", name)
}

func suma(n1 int, n2 int) int {
	return n1 + n2
}
// r = 0 at the beginning
func resta(n1, n2 int) (r int) {
	r = n1 - n2
	return
}

// numbers = slice
func sumar(numbers ...int) (result int) {
	for _, number := range numbers {
		result += number
	}
	return
}

func main() {
	imprimirNombre("Juan")
	fmt.Println("Dentro del main")
	println(suma(10, 10))
	println(resta(50, 10))

	fmt.Printf("Funcion suma: %T\n", suma)
	fmt.Printf("Funcion resta. %T\n", resta)

	fmt.Println(sumar(1, 2, 3, 4, 5, 6))
	fmt.Println(sumar())

	numbers := []int{
		25,
		10,
		20,
	}
	// the three dots indicate that we pass the values of the slice
	fmt.Println(sumar(numbers...))
}