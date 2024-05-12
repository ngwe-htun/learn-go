package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var once sync.Once

func main() {

	TestSyncOnce()
	// var wg sync.WaitGroup

	// ch := make(chan int, 3)
	// go multiply(ch)
	// a := <-ch
	// close(ch)

	// fmt.Println(a)
	// TestMutex()

	// wg.Add(1)
	// go foo(&wg)
	// wg.Wait()
	// fmt.Println("Main")
}

// func foo(wg *sync.WaitGroup) {
// 	fmt.Println("I am foo")
// 	wg.Done()
// }

// func multiply(ch chan int) {
// 	ch <- 3
// }

// Mutex - Binary semaphore
// func TestMutex() {
// 	var i int
// 	var mut sync.Mutex

// 	mut.Lock()
// 	i = i + 1
// 	mut.Unlock()

// 	fmt.Println(i)
// }

// Test sync.Once
func TestSyncOnce() {
	wg.Add(2)
	go DoSomething(&wg)
	go DoSomething(&wg)
	wg.Wait()
}

func DoSomething(wg *sync.WaitGroup) {
	once.Do(Setup)
	fmt.Println("I am doing something")
	wg.Done()
}

func Setup() {
	fmt.Println("Init")
}

// Sync package
// - Wait group
// - Mutex
// - Once
// -
