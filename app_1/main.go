package main

import (
	"fmt"
	"unicode/utf8"
)

const global_variable = 9.81
const Gravity = 9.81

func main() {
	fmt.Println("Hello there!")

	var fname = "Ahhmer"
	fmt.Println(fname)

	// Does not work in any condition
	// var num int = "244"
	// num := "244"
	// fmt.Println("Sum:", num+23)

	var name string = "Alice"
	fmt.Println(name)

	var age int = 77
	fmt.Println(age)

	fullname := "Ahmer"
	your_age := 30

	fmt.Println(fullname)
	fmt.Println(your_age)

	const pi = 3.14
	fmt.Println(pi)

	const pi1 float32 = 3.14
	fmt.Println(pi1)

	const pi2 int = 3
	fmt.Println(pi2)

	var distance = 888
	fmt.Println("Distance:", distance)

	fmt.Println("Gravity:", global_variable*3)

	var height float64 = 10
	var acceleration float64 = Gravity * height
	fmt.Println("Acceleration:", acceleration)

	var char rune = '🤡'
	fmt.Println(char)
	var char1 rune = '🙏'
	fmt.Println(char1)

	var r rune = 'A'
	fmt.Printf("Rune: %c, Type: %T, Value: %d\n", r, r, r)

	euroRune := '€'
	fmt.Printf("Rune: %c, Code Point: %U, Value: %d\n", euroRune, euroRune, euroRune)

	s := "Hello, 世界"
	fmt.Println("\nIterating over string 'Hello, 世界':")
	for index, runeValue := range s {
		fmt.Printf("Index: %d, Rune: %c, Code Point: %U, Value: %v\n", index, runeValue, runeValue, runeValue)
	}

	str := "Hello, 世界"
	runeSlice := []rune(str)
	fmt.Printf("\nString converted to a []rune slice: %v\n", runeSlice)

	byteLength := len(str)
	runeCount := utf8.RuneCountInString(str)
	fmt.Printf("String '%s' has length (bytes): %d, and count (runes): %d\n", str, byteLength, runeCount)

	var x int32 = 10
	if x > 10 {
		fmt.Println("Big")
	} else if x < 10 {
		fmt.Println("Small")
	} else {
		fmt.Println("Equal")
	}

	// Go's only loop, you can use range too
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	for x < 100 {
		x *= 2
		fmt.Println(x)
	}

	var day = "Mon"
	switch day {
	case "Mon":
		fmt.Println("Monday")
	case "Fri":
		fmt.Println("Friday")
	default:
		fmt.Println("Other day")
	}

	var result = add(9, 12)
	// fmt.Printf(result) this can not be used
	// fmt.Printf("%d\n", result) but can be used like this
	fmt.Println(result)

	var div, err = divide(12, 22.3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(div)
	}

	// Fixed size array
	arr := [3]int{1, 2, 3}
	fmt.Println(arr[0])
	fmt.Println(arr[2])

	// Dynamic array
	array := []int{1, 2, 3, 4, 5}
	array = append(array, 7)
	fmt.Println(array[1])
	fmt.Println(array[3])
	fmt.Println(array[5])

	// Map like PHP's Associative array
	m := map[string]int{"age": 25, "year": 2002}
	m["order"] = 35
	val, ok := m["age"]
	if ok {
		fmt.Println(val)
	} else {
		fmt.Printf("Value not found\n")
	}
}

// functions in go
func add(a int, b int) int {
	return a + b
}

// Common error return in Go
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}
