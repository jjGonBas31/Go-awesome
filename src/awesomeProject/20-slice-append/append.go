package main

import "fmt"

func main() {
	x := make([]byte, 4, 10)
	fmt.Println(x)

	x = []byte {'H', 'O', 'L', 'A'}
	fmt.Println(x)

	fmt.Printf("Slice x: %q \n" , x)

	for i := 0; i < len(x); i++ {
		fmt.Printf("Slice x[%d]: %q \n", i, x[i])
	}

	// x[5] = ' ' // <-- this sentence throw an error
	x = append(x, ' ') // better with append
	fmt.Println(x)

	x = append(x, 'M', 'U', 'N', 'D', 'O')

}