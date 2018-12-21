package main

import (
	"fmt"
	"strings"
)

func main() {
	cadena := "123456789"

	// ejemplo de función anónima
	cadena = strings.Map(func(r rune) rune {
		return r + 1
	}, cadena)

	fmt.Println(cadena)
}
