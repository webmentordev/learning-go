package main

import "fmt"

/*
Generic can be a function that support multiple types,
it prevents repeating same for multiple types
*/
func main() {
	var intSlice = []int{1, 2, 3}
	fmt.Println(sumSlice[int](intSlice)) // [not-needed]
	var float32Slice = []float32{1, 2, 3}
	fmt.Println(sumSlice[float32](float32Slice)) // [not-needed]

	var intSlice2 = []int{1, 2, 3}
	fmt.Println(isEmpty(intSlice2))
}

// Generic only accepts int & float
func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

// Generic accept all types
/*
It can cause issue with unsuppoted function like string
can not be added into string

func sumSliceAny[T any](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v // Casue issue
	}
	return sum
}*/

func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}
