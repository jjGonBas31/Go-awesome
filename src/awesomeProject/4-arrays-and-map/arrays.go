package main

func main(){
	// Static declaration
	var anotherArrayNumbers [20]int // The arrays are not dynamic but slices yes
	anotherArrayNumbers[0] = 100
	// Declare with Duck typing
	arrayNumbers := [...]int{2, 3, 4, 5}
	someSlice := arrayNumbers[0:2]
	// The slices pointer to the original array, reference type
	theSameSlice := someSlice
	println(someSlice[0])
	someSlice[0] = 1
	println(theSameSlice[0])

	// Use of range
	for x, y := range arrayNumbers {
		println(x, y)
	}

	// Map declaration
	monthdays := map[string]int{
		"Enero": 31,
		"Febrero": 28, // <-- the coma is required
	}
	// retrieve in a map
	println(monthdays["Enero"])

	// search in a map
	v, ok := monthdays["Diciembre"]
	println(v, ok)

	//delete in a map
	monthdays["Enero"] = 0

	// range in a map
	for x, y := range monthdays{
		println(x, y)
	}
}
