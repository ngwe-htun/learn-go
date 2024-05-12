package main

import "fmt"

// Aggregate data type

type Person struct {
	Name    string
	Age     int
	Marital int
}

func main() {

	// Using dot annotation
	var a Person
	a.Name = "Aung Aung"
	a.Age = 22
	a.Marital = 3
	fmt.Println("-- Using dot annotation --")
	fmt.Println("Age - ", a.Age)
	fmt.Println("Name - ", a.Name)
	fmt.Println("Marital - ", a.Marital)

	// Initializing
	b := new(Person)
	c := Person{Name: "Aung Aung", Age: 23, Marital: 0}
	fmt.Println("-- Initializing struct --")
	fmt.Println("Using new - ", b)
	fmt.Println("Using literal - ", c)
}
