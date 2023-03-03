package main

import (
	"fmt"
	"sync"
)

// create a function to calculate the square of a number
func calcSquare(number int, square chan int, wg *sync.WaitGroup) {
	//defer the call to the waitgroup's Done method
	defer wg.Done()

	//print the number and the square of the number
	fmt.Println("Number: ", number, " Square: ", number*number)

	//calculate the square of the number
	square <- number * number
}

func main() {
	//create a list of numbers to be squared
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//create a channel to receive the squares
	square := make(chan int)

	//create a waitgroup
	var wg sync.WaitGroup

	//add the number of goroutines to the waitgroup and start the goroutines
	for i := 0; i < len(numbers); i++ {
		wg.Add(1)
		go calcSquare(numbers[i], square, &wg)
	}

	//close the channel when all the goroutines are done
	go func() {
		wg.Wait()
		close(square)
	}()

	//calculate and print the sum of all the squared numbers received on the channel
	sum := 0
	for squareVal := range square {
		sum += squareVal
	}
	fmt.Println("Sum of squares: ", sum)
}
