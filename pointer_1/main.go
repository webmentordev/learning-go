package main

import "fmt"

func main() {
	// var p *int32 // p will hold a memory address of int32 value, this will cause runtime error when dereferencing
	// This variable does not have an address to store of a value

	// This has the address to store the value
	var p *int32 = new(int32)
	var i int32
	fmt.Printf("The value p points to is: %v", *p) // It is called dereferencing the pointer
	fmt.Printf("\nThe value of i is: %v\n", i)

	/* If no value is assigned to a variable then
	Int will get 0, string will be "", false for
	boolean e.t.c */
	*p = 10
	fmt.Println(*p)

	// Referencing memory address of a value
	p = &i
	*p = 2
	fmt.Printf("The value p points to is: %v", *p)
	fmt.Printf("\nThe value of i is: %v\n", i)

	var k int32 = 2
	i = k // This is copy value of k to i, not address

	var slice = []int32{1, 2, 3}
	var sliceCopy = slice
	sliceCopy[2] = 4
	fmt.Println(slice)
	fmt.Println(sliceCopy)
	// Value of the slice has changed for both because slices contains pointers to the underlying array

}
