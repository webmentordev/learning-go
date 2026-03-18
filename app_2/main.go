package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	Name string
	Age  int
}

// Age is int so "string" + int is not possible
// "string" + "string" is possible so conversion is required
func (p Person) Greet() string {
	return "Hi, I'm " + p.Name + " and is " + strconv.Itoa(p.Age) + " years old."
}

type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string { return "Woof!" }

type Cat struct{}

func (c Cat) Speak() string { return "Meow!" }

func makeSound(a Animal) {
	fmt.Println(a.Speak())
}

func main() {
	fmt.Println("Structs")

	p := Person{Name: "Ahmer", Age: 30}
	fmt.Println(p.Greet())

	dog := Dog{}
	makeSound(dog)

	// Pointers
	x := 42
	point := &x         // Holds the value of X in Hex
	fmt.Println(*point) // Dereference the value: print 42
	*point = 100        // Modifixes x through pointer
	fmt.Println(&point)
	fmt.Println(x)

	//===================
	// Go routines
	//===================
	go func() {
		fmt.Println("running concurrently")
	}()
	// Created channel to communicate between goroutines
	ch := make(chan int)
	go func() {
		ch <- 42 // Sends data into the channel
	}()
	valt := <-ch
	fmt.Println(valt)
}
