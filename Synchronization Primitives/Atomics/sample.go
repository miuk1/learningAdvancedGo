//sample go file for atomics

package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var counter int64

	//create goroutines to increment the counter
	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
		}()
	}

	//wait for the goroutines to finish
	time.Sleep(1 * time.Second)

	//print the counter value
	fmt.Println("Counter: ", counter)
}
