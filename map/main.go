package main

import "fmt"

// Map = Implementation of hash tables
func main() {

	// By it literals
	x := map[string]string{"name": "ngwe htun"}
	fmt.Println(x)

	// Declare, initialize
	var y map[string]string
	y = make(map[string]string)
	y["name"] = "killua"
	fmt.Println("My map - ", y)

	// Adding
	y["gender"] = "male"

	// Accessing
	fmt.Println("Name - ", y["name"])
	fmt.Println("Gender - ", y["gender"])

	// Deleteing
	delete(y, "gender")
	fmt.Println("After delete - ", y)

	// Two-value assignment
	name, p := y["name"]
	fmt.Println("Name - ", name)
	fmt.Println("Is exist - ", p)

	// Length
	fmt.Println("Length - ", len(y))

	// Looping
	for i, v := range y {
		fmt.Printf("Index - %s, value - %s \n", i, v)
	}
}
