package main

import "fmt"

func main() {
	fixedArr := [...]int32{1, 2, 3, 4, 5} // Fixed array, not slice
	fmt.Println(fixedArr)

	sliceArray := []int32{1, 2, 3, 4} // This is slice
	fmt.Println(sliceArray)
	fmt.Printf("The length is %v with capacity %v\n", len(sliceArray), cap(sliceArray))
	sliceArray = append(sliceArray, 45) // Increased elements & capacity
	fmt.Println(sliceArray)
	fmt.Printf("The length is %v with capacity %v\n", len(sliceArray), cap(sliceArray))

	var sliceArray2 []int32 = []int32{8, 9}
	sliceArray = append(sliceArray, sliceArray2...) // Append slice with slice
	fmt.Println(sliceArray)

	// Additional way to initialize the slice
	// This is fast due to pre-capacity allocation
	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice3)
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice3), cap(intSlice3))

}
