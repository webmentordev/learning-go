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
}
