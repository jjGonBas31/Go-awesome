package main

import "fmt"

func main(){

	// Alternative for first exercise

	//someText := `Por favor, introduce un número \n:`
	//fmt.Println(someText)
	//var someNumber int
	//fmt.Scanln(&someNumber)
	//someArray := [...]string{"*", "*"}
	//sliceArray := someArray[0:0]
	//
	//for i := 1; i <= someNumber; i++ {
	//	sliceArray = append(sliceArray, "*")
	//	fmt.Printf("%q \n", sliceArray)
	//}

	// Create a Go program that run the following
	array := [100]string{"A"}
	for i := 0; i < 100; i++ {
		slice := array[0:i]
		for _, y := range slice{
			print(y)
		}
		println()
		array[i] = "A"
	}

	// Fizzbuzz
	for i := 0; i <= 100; i++{
		if i % 5 == 0 && i % 3 == 0{
			println("FizzBuzz")
		} else if i % 5 == 0{
			println("Buzz")
		} else if i % 3 == 0 {
			println("Fizz")
		} else {
			println(i)
		}
	}

	// Reversed string
	fooString := "foobar"
	reversedString := [...]byte{'a', 'b', 'c', 'd', 'e', 'd'}
	j := 5
	i := 0
	for j >= 0{
		reversedString[i] = fooString[j]
		j--
		i++
	}
	for _, v := range reversedString{
		fmt.Printf("%c", v)
	}
}
