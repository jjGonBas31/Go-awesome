package main

import "fmt"

// panic se utiliza para detener errores en tiempo de ejecucion
func imprimir() {
	fmt.Println("Hola Alex!")
	//panic("Error")
	//
	//cadena := recover()
	//}





	defer func() {
		someString := recover()
		fmt.Println(someString)
	} ()
	panic("Error")







}



func main() {
	imprimir()
	fmt.Println("Hola main")
}

