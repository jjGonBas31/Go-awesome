package main

import "fmt"

func SimpleAdding(num int) int {
	if num == 0 {
		return 0
	}
	return num + SimpleAdding(num - 1)
}

func main() {
	fmt.Println(SimpleAdding(140))
}