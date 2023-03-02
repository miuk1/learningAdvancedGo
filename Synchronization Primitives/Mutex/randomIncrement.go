//Problem: You have two counters, c1 and c2, that are initially zero. You need to run n concurrent goroutines where each goroutine randomly increments one of the counters k times (k is a random integer between 1 and 10). Once all goroutines have completed, you need to print the final values of the counters.

// Path: Mutex\randomIncrement.go

package main

import (
	"fmt"
	"math/rand" // for random number generation
	"sync"      // for mutex
	"time"      // for time
)

var (
	// counter 1 and counter 2 is intialized to 0
	count1 = 0        // counter 1
	count2 = 0        // counter 2
	mutex  sync.Mutex // mutex to synchronize access to the counters
)

func main() {
	//set the seed value for the random number generator
	rand.Seed(time.Now().UnixNano())

	//create 10 goroutines that randomly increment one of the counters
	for i := 0; i < 10; i++ {
		go func() {
			//genrate a random number between 1 and 10
			num := rand.Intn(10) + 1
			mutex.Lock()

			if rand.Intn(2) == 0 {
				count1 += num
			} else {
				count2 += num
			}
			mutex.Unlock()
		}()
	}

	//wait for the goroutines to finish
	//show a message to user and wait for a key to be pressed
	fmt.Println("Press any key to continue...")
	fmt.Scanln()

	//print the values of the counters
	fmt.Println("Count1 is: ", count1)
	fmt.Println("Count2 is: ", count2)
}
