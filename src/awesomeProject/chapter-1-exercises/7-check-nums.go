package main

/*Have the function CheckNums(num1,num2) take both parameters being passed and return the string true if num2 is
greater than num1, otherwise return the string false. If the parameter values are equal to each other
then return the string -1. */

import "fmt"

func CheckNums(num1 int, num2 int) (result string) {
	if num2 == num1 {
		result = "-1"
	} else if num2 > num1 {
		result = "true"
	} else {
		result = "false"
	}
	return
}
func main() {
	fmt.Println(CheckNums(4, 8))
}