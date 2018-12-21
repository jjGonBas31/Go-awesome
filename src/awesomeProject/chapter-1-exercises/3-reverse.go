package main

import "fmt"

func FirstReverse(str string) (reversed string) {
	for i := len(str) - 1; i >= 0; i--{
		reversed += string(str[i])
	}
	return
}

func main (){
	fmt.Println(FirstReverse("foobar"))
}
