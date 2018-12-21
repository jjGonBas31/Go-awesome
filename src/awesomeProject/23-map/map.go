package main

import "fmt"

func main() {
	// declare a map
	x :=  make(map[string]string)
	fmt.Println(x)

	// another form more efficiently
	y := make(map[string]string, 7)
	fmt.Println(y)

	// add values
	x["name"] = "Alejandro"
	x["age"] = "29"
	fmt.Println(x)

	// access to a value
	fmt.Println(x["name"])
	fmt.Println(x["age"])

	ages := map[string]int{
		"Ana": 55,
		"Manuel": 10,
	}

	fmt.Println(ages)

	days := map[int]string{
		1: "Monday",
		2: "Tuesday",
	}

	delete(days, 1)

	//  if some key not exist the map return 0
	fmt.Println(days[3])
	// or delete
	delete(days, 100)

	// Assign 1 to key "Pedro", the default value is 0
	ages["Pedro"]++
	// add values to some entry
	ages["Pedro"] =+ 2

	// map return two values when we make a consult
	// in this consult ok receive a boolean value with true or false depending of the existence of the key, value pair
	// in the map
	if age, ok := ages["Manuel"]; ok {
		if age < 18 {
			fmt.Println("Es menor de edad")
		} else {
			fmt.Println("Es mayor de edad")
		}
	} else {
		fmt.Println("El valor no existe")
	}

	//len of map
	fmt.Println(len(ages))

	// the elements of the map aren't variables
	// we cannot obtain the pointer
	// puntero := &ages["Carlos"]
}
