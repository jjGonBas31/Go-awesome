package main

import "fmt"

//noinspection GoUnhandledErrorResult
func main() {
	// counter de números impares
	header := `
	****************************
	Contador de números impares
	****************************`
	fmt.Println(header)

	fmt.Printf("Digita el primer número")
	var numberOne, numberTwo = 0, 0
	fmt.Scanln(&numberOne)
	fmt.Println("Digita el número dos")
	fmt.Scanln(&numberTwo)
	var counter int
	for i := numberOne; i <= numberTwo; i++ {
		if i % 2 != 0 {
			counter++
			fmt.Printf("%d es un numero impar\n", i)
		}
	}

	header = `
	****************************
	Resultado del conteo
	****************************`
	fmt.Println(header)
	fmt.Printf("Entre el %d y %d hay %d numeros impares", numberOne, numberTwo, counter)

}
