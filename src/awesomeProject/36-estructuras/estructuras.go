package main

import (
	"fmt"
)

// kind off de un objeto en cualquier otro lenguaje
type Persona struct {
	Nombre string
	Edad int
}

func Older (p1, p2 Persona) (Persona, int){
	if p1.Edad > p2.Edad {
		return p1, p1.Edad - p2.Edad
	}
	return p2, p2.Edad  - p1.Edad
}

func main() {
	// forma #1 de como declarar un objeto y asignar sus propiedades
	var p Persona
	p.Nombre = "Alejandro"
	p.Edad = 18
	fmt.Println("Estructura p de tipo persona ", p)

	// forma #2 de como declarar un objeto y asignar sus propiedades
	p2 := Persona{Nombre: "Rafael", Edad: 21}
	fmt.Println("Nombre de p2", p2.Nombre)
	fmt.Println("Edad de p2", p2.Edad)

	// forma #3 de como declarar un objeto y asignar sus propiedades
	p3 := Persona{"Yamilet", 17}
	fmt.Println(p3)

	tom := Persona{"Tom", 18}
	bob := Persona{"Bob", 32}
	//paul := Person {"Paul", 100}

	tbOlder, tbDiff := Older(tom, bob)
	//tpOlder, tpDiff := Older(tom, paul)
	//bpOlder, bpDiff := Older(bob, paul)

	fmt.Printf("Entre %s y %s, %s es mayor por %d años\n", tom.Nombre, bob.Nombre, tbOlder, tbDiff)


}
