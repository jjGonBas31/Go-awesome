package main

import "fmt"

// las funciones también son un tipo
func print(cadena string){
	fmt.Println(cadena)
}

func impPrint(cadena string){
	fmt.Println(cadena)
}

// ejemplo de funciones con un parámetro tipo función
func printFour(fprint func(string)){
	fprint("Hola mundo desde Print Four")
}

func incrementar() func() int {
	i := 0
	return func() (r int) {
		r = i
		i += 2
		return
	}
}

func main() {
	cadena := "Hola mundo"

	// asignación de la función print a la variable imprimir
	imprimir := print

	// función dentro de otra función
	imprimir(cadena)

	// closure
	// cadena está disponible debido al scope
	anotherPrint := func() {
		fmt.Println(cadena)
	}

	// llamada a anotherPrint
	anotherPrint()

	// funciones con la misma firma se pueden reasignar
	imprimir = impPrint

	// esto causará un error por no tener la misma firma
	//imprimir = anotherPrint

	printFour(imprimir)

	inc := incrementar()
	// Ámbito de la variable
	fmt.Println("Valor de i", inc())
	fmt.Println("Valor de i", inc())
	fmt.Println("Valor de i", inc())


}
