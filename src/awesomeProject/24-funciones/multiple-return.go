package main

import "fmt"

// multiple return
func multlipicar(numero int) (n1, n2, n3 int){
	n1 = numero * 10
	n2 = numero * 20
	n3 = numero * 30
	return
}

func main() {
	fmt.Println(multlipicar(5))
}
