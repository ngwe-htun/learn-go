package main

import (
	"ngwehtun/learn-go/class/ens"
)

type MyInt int

type Person struct {
	Name  string
	Email string
}

func main() {

	// Receiver type
	// a := MyInt(3)
	// doubled := a.Double()
	// fmt.Println("Doubled - ", doubled)

	// Using struct
	// p := &Person{Name: "Ngwe Htun"}
	// fmt.Println(p.GetName())

	// Publis with capital
	//fmt.Println(ens.IsEns())
	// Private with lowercase
	// cannot call ens.isenv()

	cat := ens.Cat{}
	cat.SetName("Jhon")
	cat.PrintName()

}

// Type should be in same packages
// Receiver
func (m MyInt) Double() int {
	return int(m * 2)
}

func (p *Person) GetName() string {
	return p.Name
}
