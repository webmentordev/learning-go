package main

import "fmt"

func main() {
	// Int is 4 bytes so Go allocated contiguous memory
	// Array are fixed size
	var fixedArray [3]int32 // All values are 0
	fmt.Println(fixedArray[0])
	fmt.Println(fixedArray[1:3]) // Access two elements at a time

	var intArray [3]int32
	intArray[1] = 123 // Update value at an index
	fmt.Println(intArray[1])
	fmt.Println(&intArray[2]) // Print memory location using &

	var initArray [3]int32 = [3]int32{1, 2, 3}       //Pre intialize array
	init2Array := [...]int32{1, 2, 3, 4, 5, 6, 7, 8} // Auto infer type and size
	fmt.Println(initArray)
	fmt.Println(init2Array)
}
