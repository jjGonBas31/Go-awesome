package main

import "fmt"

func incrementar(numero *int){
	*numero++
	fmt.Println("Valor de número en la función", *numero)
}

func main() {
	/*a := 25
	fmt.Println("Valor de a: ", a)
	// em ampersand indica la dirección de memoria
	fmt.Println("Dirección de a: ", &a)

	// b ahora almacena el puntero
	b := &a
	fmt.Println(b)

	// el asterisco antes de la variable indica que ahora queremos el valor del apuntador
	fmt.Println(*b)

	// esto marcará error
	//&b = 25

	// esta operación es válida
	*b = 90
	fmt.Println(a) // a ahora tiene el valor de 90 ya que modificamos el valor de la referencia del puntero

	if b != nil {
		fmt.Println("b es diferente de nil")
	}

	c := &a

	// los punteros son comparables
	if b == c {
		fmt.Println("b y c son iguales")
	}

	// creación de un puntero con new()
	d := new(int)
	fmt.Println("Dirección de d: ", d) // dirección del puntero
	fmt.Println("Valor de d: ", *d) // 0

	d = b

	// ahora todos los punteros apuntan a la misma dirección
	fmt.Println("Valor de a: ", a)
	fmt.Println("Valor de b: ", b)
	fmt.Println("Valor de c: ", c)
	fmt.Println("Valor de d: ", d)*/

	numero := 2
	fmt.Println("Numero antes de incrementar: ", numero)

	// funciones con punteros como parámetros
	incrementar(&numero)
	fmt.Println("Número después de incrementar: ", numero)


}
