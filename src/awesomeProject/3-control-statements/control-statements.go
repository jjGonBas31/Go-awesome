package main

func main(){

	myFunc()
	sum := 0

	// for like java
	for  i := 0; i < 10; i++ {
		sum += 1
	}

	// for like while
	number := 0
	for number < 10 {
		number++
	}

	// label for cycles
	J: for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			if i > 5 {
				break J // specify cycle to break
			}
			println(i)
		}

	}
}

// goto statement, use wisely
func myFunc(){ // example of an eternal cycle
	i := 0
HERE: // <-- first word on a line ending with a colon is a label
	println(i)
	i++
	goto HERE // <-- Jump
}
