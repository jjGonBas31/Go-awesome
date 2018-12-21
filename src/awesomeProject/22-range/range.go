package main

import "fmt"

func main() {
	nombres := []string {
		"Alejandro",
		"Maria",
		"Pedro",
		"Carlos",
	}

	// use of range
	for index, nombre := range nombres {
		fmt.Printf("El nombre %q esta en el index %d \n", nombre, index)
	}

	// unused index
	for _, nombre := range nombres {
		fmt.Printf(nombre)
	}

	cadena := "Hola mi gente bonita"
	for index, letra := range cadena {
		fmt.Printf("La letra %q está en el index %d \n", letra, index)
	}


}
