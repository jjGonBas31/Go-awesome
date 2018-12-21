package main

import "fmt"

// para crear nuevos tipos se tiene que partir de un tipo ya existente
type Dinero int // creación del tipo Dinero

// recibidor -> (d Dinero)
func (d Dinero) String() string {
	return fmt.Sprintf("$%d", d) // -> Funcionalidad de declarar nuestros propios tipos
}

func main() {
	// asignación del tipo Dinero para una variable
	var sueldo Dinero
	sueldo = 25000
	fmt.Println("Sueldo: ", sueldo)

	aumento := 10000
	sueldo += Dinero(aumento) // -> Necesitamos hacer una conversión de tipos
	fmt.Println("Sueldo + Aumento ", sueldo)

}


