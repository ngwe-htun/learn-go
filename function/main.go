package main

import "fmt"

func main() {

	//---------------------
	// variable arguments -
	//---------------------
	// funcVariableArgs(1, 2, 3)
	// funcVariableArgs([]int{1, 2, 3}...)

	//--------
	// Defer -
	//--------
	// testDefer()
}

func twoReturn(x int) (int, int) {
	return x, x
}

// Pass by value
func passByValue(v int) int {
	return v + 1
}

// Pass by reference
func passByRef(v *int) {
	*v = *v + 1
}

// Pass array as argument
func passArray(v []int) []int {
	fmt.Println(v)
	return v
}

// Function naming
// Parameter naming
// Function cohesion - only one operation
// Make them simpler - reduce parameter

// First-class variable
// Variable as a function
func funcAsAvariable() {

	var test func(int) int
	test = func(i int) int { return i }
	fmt.Println(test(9))
}

// Function as argument
func funcAsArgument(v func(int) int) {
	fmt.Println("result - ", v(1))
}

// Anonymous function
func anonymousFunc() {
	a := func() {
		fmt.Println("I am just anonymous function")
	}
	a()
}

// Variable arguments
func funcVariableArgs(values ...int) {
	fmt.Println("Number of arguments - ", len(values))
	for i, v := range values {
		fmt.Printf("Argument %d - %d\n", i, v)
	}
}

// Defer
func testDefer() {
	defer fmt.Println("World")
	fmt.Println("Hello ")
}
