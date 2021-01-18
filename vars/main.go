package main

import "fmt"

func main() {
	// Main type
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// float32 float64

	var name string = "Sumit"
	var age int = 56

	// shorthand variable declaration

	name2 := "Saroj"
	size := 1.3

	//
	name2, email := "Saroj", "saroj@gmail.com"

	var isCool = true

	fmt.Println(name, name2, age, isCool, size, email)

	// Get data type
	fmt.Printf("%T\n", size)
}
