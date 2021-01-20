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

	// emails := map[string]string{"Bob": "bob@gmail.com", "mike": "mike@gmail.com"}

	// fmt.Println(emails)

	// delete from map
	// delete(emails, "Bob")

	// fmt.Println(emails)

	// new  - only allocates - no init of memoey\
	// make - allocate and init

	score := make(map[string]int)
	score["sumit"] = 89

	score["Josh"] = 98
	score["mesh"] = 56
	score["kesh"] = 7

	fmt.Println(score)
	getJoshScore := score["Josh"]
	fmt.Println(getJoshScore)

	delete(score, "kesh")

	fmt.Println("Updated score", score)

	for k, v := range score {
		fmt.Printf("Score of %v is %v\n", k, v)
	}
}
