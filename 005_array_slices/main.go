package main

import "fmt"

func main() {
	// Arrays
	var fruitArray [2]string

	// Assign value
	fruitArray[0] = "Apple"
	fruitArray[1] = "Orange"

	// declare and assign

	gameArray := [2]string{"Football", "cricket"}

	fmt.Println(fruitArray, gameArray)

	// with out specifying array size
	fruitSlice := []string{"Apple", "orange", "grape", "blueberry"}

	fmt.Println(len(fruitSlice))

	fmt.Println(fruitSlice[1:2])
}
