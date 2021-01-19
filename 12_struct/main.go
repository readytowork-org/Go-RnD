package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type person struct {
	firstName string
	age       int
}

// Greeting method (value reciever)

func (p person) greet() string {
	return "Hello, my name is " + p.firstName + " And age is :" + strconv.Itoa(p.age)
}

// Define car struct
type car struct {
	brand string
	model string
}

// a demo method
func (c car) getBrand() string {
	return "Brand name is: " + c.brand
}

func main() {
	// Init person using struct
	person1 := person{"Sumit", 24}
	car1 := car{"Ford", "m1"}

	// fmt.Println(person1)
	// fmt.Println("Name: ", person1.firstName)

	fmt.Println(person1.greet())

	fmt.Println(car1.getBrand())

}
