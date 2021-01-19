package main

import "fmt"

func main() {
	x := 15
	y := 10

	if x < y {
		fmt.Printf("%d is smaller than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}
}
