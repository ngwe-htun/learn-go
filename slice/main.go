package main

import "fmt";

func main () {

	// array of string
	x := [...]string{"a", "b", "c", "d", "e", "f", "g"};

	// slice
	s1 := x[0:2];
	s2 := []string{"x", "y", "z"}

	// using make()
	// make slice without define capacity
	// s3 := make([]int, 10);
	// make sliece with capacity
	s4 := make([]int, 10, 15);

	// append
	s5 := []int{1,2,3}
	s5 = append(s5, 4);
	fmt.Println("Append result - ", s5);

	// Printing
	fmt.Println("slice 1 - ", s1);
	fmt.Println("length of s1 - ", len(s1));
	fmt.Println("capacity of s1 - ", cap(s1));
	fmt.Println("slice 2 - ", s2);

	fmt.Printf("s4 - length - %d, capacity - %d \n", len(s4), cap(s4));
}