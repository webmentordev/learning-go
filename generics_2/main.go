package main

import "fmt"

type gasEngine struct {
	gallons float32
	mpg     float32
}

type electicEngine struct {
	kwh   float32
	mpkwh float32
}

// Generic struct
type car[T gasEngine | electicEngine] struct {
	carMake  string
	carModel string
	engine   T
}

func main() {
	var gasCar = car[gasEngine]{
		carMake:  "Honda",
		carModel: "Civic",
		engine: gasEngine{
			gallons: 12.4,
			mpg:     40,
		},
	}
	fmt.Println(gasCar)

	var electicCar = car[electicEngine]{
		carMake:  "Tesla",
		carModel: "Model 3",
		engine: electicEngine{
			kwh:   57.5,
			mpkwh: 4.17,
		},
	}
	fmt.Println(electicCar)
}
