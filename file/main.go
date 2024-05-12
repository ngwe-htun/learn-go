package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {

	// Reading small file
	d, _ := os.ReadFile("/workspaces/learn-go/file/test.txt")
	fmt.Printf("Read data - %s \n", d)

	// writing small file
	os.WriteFile("/workspaces/learn-go/file/test.txt", []byte("I write new content"), fs.ModeAppend)

	// Read wit limitation
	f, _ := os.Open("/workspaces/learn-go/file/test.txt")
	limit := make([]byte, 1)
	f.Read(limit)
	f.Close()

	fmt.Printf("Read - %s \n", limit)
}
