package main

import "fmt"

func main() {

	// Map is a key value pair
	// Define map
	// emails := make(map[string]string)

	// Assign key value

	// emails["Bob"] = "Bob@gmail.com"
	// emails["Sharon"] = "sharon@gmail.com"
	// emails["mike"] = "mike@gmail.com"

	// Declare map and add kv

	emails := map[string]string{"Bob": "bob@gmail.com", "mike": "mike@gmail.com"}

	fmt.Println(emails)

	// delete from map
	// delete(emails, "Bob")

	fmt.Println(emails)

}
