package main

import (
	"fmt"
)

type Person struct {
	// Person tiene 2 campos
	Nombre string
	Apellido string
}

type Student struct {
	// anonymous type
	Person // Student ya tiene las propiedades y los métodos de la estructura persona, ventaja de un tipo anónimo
	Carrera string
}

type Profesor struct {
	// another anonymous type
	Student
	Class string
}

func main (){
	alejandro := Student{
		// declaración de un tipo anónimo
		Person{"Alejandro", "Gutiérrez"},
		"Informática",
	}

	fmt.Println(alejandro)

	// accedes a las propiedades directamente, no importa de dónde viene
	fmt.Println(alejandro.Nombre)
	fmt.Println(alejandro.Apellido)
	fmt.Println(alejandro.Carrera)

	pedro := Profesor{
		Student {
			 Person {
				"Pedro",
				"González",
			},
			"Informática",
		},
		"Contabilidad",
	}

	// accediendo a propiedades con el mismo nombre
	fmt.Println(pedro.Student.Carrera)
	fmt.Println(pedro.Class)

}
