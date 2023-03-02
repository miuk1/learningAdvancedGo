package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch := make(chan int)

	// Create two goroutines to fill ch1 and ch2 with some data
	go func() {
		for i := 1; i <= 5; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	go func() {
		for i := 6; i <= 10; i++ {
			ch2 <- i
		}
		close(ch2)
	}()

	// Create a goroutine to merge the results of ch1 and ch2 into ch
	go func() {
		sum := 0
		done := make(chan bool)

		go func() {
			for i := range ch1 {
				sum += i
			}
			done <- true
		}()

		go func() {
			for i := range ch2 {
				sum += i
			}
			done <- true
		}()

		<-done
		<-done

		ch <- sum
		close(ch)
	}()

	// Read the results from ch
	result := <-ch
	fmt.Println(result)
}
