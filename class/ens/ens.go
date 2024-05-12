package ens

import "fmt"

var isEns = true

type Cat struct {
	name string
}

func IsEns() bool {
	return isEns
}

// Private func
// func isenv() bool {
// 	return isEns
// }

func (c *Cat) SetName(name string) {
	c.name = name
}

func (c *Cat) PrintName() {
	fmt.Println("Cat name - ", c.name)
}
