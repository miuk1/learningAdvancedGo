/*
Write a Go program that calculates the sum of squares of a list of numbers concurrently.
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	//prompt user to enter numbers
	fmt.Println("Enter a number or press q to quit: ")
	var input string
	var numbers []int

	for {
		fmt.Print(">")
		fmt.Scanln(&input)
		if input == "q" {
			break
		} else {
			number, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Invalid input, please enter a number or press q to quit")
			} else {
				numbers = append(numbers, number)
			}
		}
	}

	ch := make(chan int)

	//creating a gorutine that runs calculateSumOfSquares function
	go calculateSumOfSquares(numbers, ch)

	//printint the sum of squares
	fmt.Println("Sum of squares of numbers is: ", <-ch)

}

// calculateSumOfSquares calculates the sum of squares of a list of numbers
func calculateSumOfSquares(numbers []int, ch chan int) {
	sum := 0
	for _, v := range numbers {
		sum += v * v
	}
	ch <- sum
}
