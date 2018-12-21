package main

import (
	"fmt"
	"os"
)

func primera(){
	fmt.Println("Primera")
}

func segunda(){
	fmt.Println("Segunda")
}

func main() {
	// defer se reserva para ejecutar la función hasta el final del código
	// defer siempre se ejecutará
	defer primera()
	segunda()

	f, err := os.Open("text.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	data := make([]byte, 175)
	c, err := f.Read(data)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Cantidad de bytes leídos: %d\n Texto leído n%q \nerror: %v", c, data, err)
}
