package main

import (
	"fmt"
	"strings"
)

func LetterCapitalize(str string) (capStr string) {
	capStr = strings.ToUpper(string(str[0]))

	for i := 1; i < len(str); i++ {
		if str[i] == 32 {
			capStr += string(str[i])
			capStr += strings.ToUpper(string(str[i + 1]))
			i += 1
		} else {
			capStr += string(str[i])
		}
	}
	return
}

func main() {
	fmt.Println(LetterCapitalize("hello world"))
}
