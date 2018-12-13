package main

import "fmt"

func main() {
	// declare slices
	var j []int
	fmt.Println("Slice j: ", j)

	x := []int{1, 2, 3}
	fmt.Println(x)

	y := make([]int, 5)
	fmt.Println("Slice y: ", y)
	fmt.Println("Longitud de y: ", len(y))
	fmt.Println("Capacidad de y ", cap(y)) // <-- isn't the same of len

	// make an slice (type, len, capacity)
	k := make([]int, 5, 10)
	fmt.Println("Slice k: ", k)
	fmt.Println("Longitud de k: ", len(k))
	fmt.Println("Capacidad de k: ", cap(k)) // <-- isn't the same of len

	// declare an array with 10 int elements
	var ar = [10]int {0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	var a, b []int
	a = ar[3:5]
	b = ar[5:8]

	fmt.Println(a, cap(b))

}
