package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

type dieselEngine struct {
	mpg       uint8
	gallons   uint8
	ownerName owner
	// owner // We can also use owner only
}
type owner struct {
	name string
}

// ##############
// Interfaces
// ##############
type engine interface {
	milesLeft() uint8
}

func main() {
	// Struct are custom data types

	// var myEngine gasEngine = gasEngine{mpg: 25, gallons: 15}
	// We can omit the field names
	var myEngine gasEngine = gasEngine{25, 15}
	fmt.Println(myEngine.mpg, myEngine.gallons)

	var newEngine gasEngine
	newEngine.gallons = 12
	newEngine.mpg = 30
	fmt.Println(newEngine.mpg, newEngine.gallons)

	var dieselEngine dieselEngine = dieselEngine{14, 23, owner{"Ahmer"}}
	fmt.Println(dieselEngine.mpg, dieselEngine.gallons, dieselEngine.ownerName.name) // If we use just owner then dieselEngine.name will be used

	// Anonymous struct type, not reuseable
	var gasEngine2 = struct {
		mpg     uint8
		gallons uint8
	}{12, 15}
	fmt.Println(gasEngine2.mpg, gasEngine2.gallons)

	// Using method
	var myEngine3 gasEngine = gasEngine{25, 15}
	fmt.Printf("Total miles left in tank: %v\n", myEngine3.milesLeft())

	// Using generic methods
	var gasEngine4 gasEngine = gasEngine{12, 25}
	canMakeIt(gasEngine4, 50)
	var electricEngine electricEngine = electricEngine{12, 25}
	canMakeIt(electricEngine, 50)
}

// functions but as method
func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}
func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

// Generic method
func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("Need to fuel up first!")
	}
}
