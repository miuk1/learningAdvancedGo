package main

import (
	"fmt"
	"sync"
)

var (
	count int
	//create a Mutex to synchronize access to the count variable
	mutex sync.Mutex
)

func main() {
	//create 10 goroutines that each increment the count variable
	for i := 0; i < 10; i++ {
		go func() {
			//lock the mutex
			mutex.Lock()
			//increment the count variable
			count++
			//unlock the mutex
			mutex.Unlock()
		}()
	}

	//wait for the goroutines to finish
	//show a message to user and wait for a key to be pressed
	fmt.Println("Press any key to continue...")
	fmt.Scanln()

	//print the value of the count variable
	fmt.Println("Count is: ", count)
}
