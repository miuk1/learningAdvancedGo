//write a program that creates 10 goroutines that each write 10 random numbers to a channel. Then pull the numbers off the channel and print them.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//set the seed value for the random number generator
	rand.Seed(time.Now().UnixNano())

	//create a channel to receive the random numbers
	ch := make(chan int)

	//create 10 goroutines that each write 10 random numbers to the channel
	for i := 0; i < 10; i++ {
		go func() {
			//generate a random number between 1 and 100
			num := rand.Intn(100) + 1

			//send the number to the channel
			ch <- num
		}()
	}

	//pull the numbers off the channel and calculate the sum
	sum := 0
	for i := 0; i < 10; i++ {
		num := <-ch
		sum += num
		fmt.Println("Random number: ", num)
	}
	// Print a separator between the random numbers and the sum
	fmt.Println("--------------------")
	//print the sum
	fmt.Println("Sum of random numbers is: ", sum)
}
