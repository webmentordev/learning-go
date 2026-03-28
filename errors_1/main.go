package main

import (
	"errors"
	"fmt"
)

func main() {
	var nominator int = 30
	var denominator int = 0
	var result, remainder, err = intDivider(nominator, denominator)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of integer division is %v.\n", result)
	} else {
		fmt.Printf("The result of integer division is %v and reminder is %v\n", result, remainder)
	}
}

func intDivider(nominator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Can not divide by 0")
		return 0, 0, err
	}

	var result int = nominator / denominator
	var remainder int = nominator % denominator

	return result, remainder, err
}
