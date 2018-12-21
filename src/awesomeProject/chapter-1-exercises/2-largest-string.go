package main

/* Have the function LongestWord(sen) take the sen parameter being passed and
return the largest word in the string. If there are two or more words that
are the same length, return the first word from the string with that length.
Ignore punctuation and assume sen will not be empty. */

import (
	"fmt"
	"strings"
)

func main() {
	someText := "fun&!! time"
	fmt.Println(LongestWord(someText))
}

func LongestWord(sen string) string {

	replacer := strings.NewReplacer("!", "", "&", "", "[", "", "]", "")
	sen = replacer.Replace(sen)

	words := strings.Split(sen, " ")
	maxIndex := 0

	for i := 1; i < len(words); i++{
		if len(words[maxIndex]) < len(words[i]){
			maxIndex = i
		}
	}
	return words[maxIndex]
}
