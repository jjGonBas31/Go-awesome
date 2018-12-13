package main

import "fmt"

func main() {
	// like Java declaration
	var x [3]int
	fmt.Println(x)

	// Go initialize any array with 0
	var k [3][2]float64
	fmt.Println(k)

	// assign a value
	x[1] = 25
	fmt.Println(x[1])

	// directly assign and declaration
	y := [5]int {1, 2, 3, 4, 5}
	fmt.Println(y)

	i := [...]int{
		1,
		2,
		// 3,
		4,

	}
	fmt.Println(i)

	// use of len
	fmt.Println(len(i))

	a := [3]int {1, 2, 3}
	b := [...]int {1, 2, 3}
	c := [3]int {1, 2, 3}

	//Compare some arrays
	fmt.Println(a == b)
	fmt.Println(a == c)

}
