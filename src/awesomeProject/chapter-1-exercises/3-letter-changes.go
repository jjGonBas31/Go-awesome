package main

import (
	"fmt"
	"strings"
)

func LetterChanges(str string) (result string) {
	for _, char := range str {
		if (char > 64) && (char < 91) || (char > 96) && (char < 123) {
			result += string(char + 1)
		} else {
			result += string(char)
		}
	}
	replace := strings.NewReplacer("a", "A", "e", "E", "i", "I", "o", "O", "u", "U")
	result = replace.Replace(result)
	return
}

func main() {
	fmt.Println(LetterChanges("a confusing /:sentence:/[ this is not!!!!!!!~"))
}
