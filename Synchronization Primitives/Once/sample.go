package main

import (
	"fmt"
	"sync"
)

var (
	value       int
	initialized bool
	mutex       sync.Mutex
)

func initialize() {
	// perfom the initialization
	fmt.Println("Performing initialization")
	value = 36
	initialized = true
}

func getValue() int {
	mutex.Lock()
	defer mutex.Unlock()

	// use Once to ensure initialization is only performed once
	var once sync.Once
	once.Do(initialize)

	if initialized {
		return value
	} else {
		initialized = true
		return 0
	}
}

func main() {
	// access the value multiple times
	for i := 0; i < 10; i++ {
		val := getValue()
		//print the counter and the value
		fmt.Println("Counter: ", i, " Value: ", val)
	}

	// wait for the user to press a key
	fmt.Println("Press any key to continue...")
	fmt.Scanln()
}
