package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Do some work
	fmt.Printf("Worker %d starting\n", id)
	// simulate some work with a sleep
	for i := 0; i < 100000000; i++ {
		// do nothing
	}
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	// Start 3 worker goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// Wait for all workers to complete
	wg.Wait()

	fmt.Println("All workers completed")
}
