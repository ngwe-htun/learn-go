package main

import (
	"fmt"
)

func main() {
	letSpeak(Cat{})
	letSpeak(Dog{})
	LetDoAnything([]int{1, 2, 3})
}

type Speaker interface {
	Speak()
}

type Cat struct{}
type Dog struct{}

// Cat
func (c Cat) Speak() {
	fmt.Println("Meow !")
}

// Dog
func (t Dog) Speak() {
	fmt.Println("Woof !")
}

func letSpeak(animal Speaker) {
	animal.Speak()

	// Check concrete type
	// Type assertion
	dog, ok := animal.(Dog)
	if ok {
		fmt.Println("Got dog - ", dog)
	}
	switch sh := animal.(type) {
	case Cat:
		fmt.Println(sh)
	case Dog:
		fmt.Println(sh)
	}
}

// Any interface
func LetDoAnything(val interface{}) {
	fmt.Println("This is the value you set - ", val)
}
