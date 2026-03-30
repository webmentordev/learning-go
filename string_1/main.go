package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "résumé"
	var indexed = myString[0] // Reads Index of underlying byte array
	// This gets first half of the UTF8 encoding.
	// To fix this, use rune type
	fmt.Printf("%v, %T", indexed, indexed)
	for i, v := range myString {
		// Range return full UTF-8 encoding
		fmt.Println(i, v)
	}

	var myString2 = []rune("résumé")
	var indexed_2 = myString2[1]
	fmt.Printf("%v, %T", indexed_2, indexed_2)

	var strSlice = []string{"a", "b", "c", "d"}
	var strBuilder strings.Builder // Builder is more efficient than using + to concatinate
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Printf("\n%v", catStr)
}
