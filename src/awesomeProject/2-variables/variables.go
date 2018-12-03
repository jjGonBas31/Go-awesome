package main

import "fmt"

func main(){ // Variables declaration
	var someNumber int = 15 // the type always goes after the name of the variable /
	someBool := false // duck notation

	var (
		anotherNumber int // Group declaration of variables
		awesomeNumber int
	)

	// a, b := 15, 20 // Parallel declaration of variables
	// unused statements causes errors

	const (
		a = iota // iota value makes an automatic plus 1 every declaration
		b = iota
	)

	var c complex64 = 5 + 5i

	println("Esta es una variable int: ", someNumber, " y esto es un bool: ", someBool)
	println(anotherNumber, ", ", awesomeNumber)
	println(a, " es la primera iota mientras que ", b, " es la segunda")
	fmt.Printf("Este es un número complejo %v", c)
}
