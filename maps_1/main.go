package main

import "fmt"

func main() {
	// Map is like Key:Value pair, like PHP's associative array
	var myMap map[string]uint8 = make(map[string]uint8) // string is the key type, uint8 is value type
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Ahmer": 14, "Ali": 25, "Shahzad": 23}
	fmt.Println(myMap2)
	fmt.Println(myMap2["Ahmer"])
	fmt.Println(myMap2["Kaleem"]) // Will return 0 becasue value ws not found

	var age, ok = myMap2["Abdullah"] // Ok will be boolean, true if value exist
	if ok {
		fmt.Printf("The age is %v\n", age)
	} else {
		println("Invalid Name")
	}

	delete(myMap2, "Ali") // Value will not be deleted by reference if not found. No error
	fmt.Println(myMap2)

	// For[Range] Loop Map, Output will not be in order when iterating map.
	for name, age := range myMap2 {
		fmt.Printf("Name: %v and age: %v\n", name, age)
	}

	// For[Range] Loop Array/Slice
	intLoopArray := []int32{1, 2, 3, 4, 5, 6, 7, 8}
	for i, v := range intLoopArray {
		fmt.Printf("Index: %v has value: %v\n", i, v)
	}
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
