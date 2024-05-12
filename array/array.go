package main;

import "fmt"

func main() {

	// Initialize first and assign
	var x [3]int;
	x[0] = 1;

	// Initialize and assign at the same time
	var y [4]int = [4]int{1,2,3,4};

	// Without specifying size
	z := [...]int{0, 9, 8, 7};

	
	// Loop through array
	for i, v := range x {
		fmt.Printf("x, index - %d, val - %d \n", i, v);
	}

	// size
	fmt.Println("size of x - ", len(x));

	// print out
	fmt.Println("y - ", y);
	fmt.Println("x - ", x);
	fmt.Println("z - ", z);
}

