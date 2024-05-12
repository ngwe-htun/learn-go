package main

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Model int
	Color string
}

func main() {

	// Json marshalling
	toyota := Car{Model: 2012, Color: "red"}
	a, _ := json.Marshal(toyota)

	fmt.Printf("Json marshal result - %s\n", a)

	// Json unmarshalling
	var ctoyata Car
	json.Unmarshal(a, &ctoyata)
	fmt.Println("Unmarshal result - ", ctoyata)
}
