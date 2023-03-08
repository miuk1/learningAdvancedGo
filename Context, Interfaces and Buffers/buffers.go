//Demonstrate buffers in Go Channels

package main

import (
	"fmt"
)

func main() {
	buffer := make(chan string, 2)
	buffer <- "Hello"
	buffer <- "World"

	//use select statement, print first buffer message and send another message, this will block because there is space in the buffer
	select {
	case buffer <- "Buffered":
		fmt.Println("Message sent to buffer")
	default:
		fmt.Println("No message sent to buffer")
	}

	//use select statement, print both buffer messages and display buffer empty message, this will not block because there is no space in the buffer
	select {
	case msg1 := <-buffer:
		fmt.Println("Recieved message from buffer: ", msg1)
	default:
		fmt.Println("Buffer empty")
	}

	select {
	case msg2 := <-buffer:
		fmt.Println("Recieved message from buffer: ", msg2)
	default:
		fmt.Println("Buffer empty")
	}

	select {
	case msg3 := <-buffer:
		fmt.Println("Recieved message from buffer: ", msg3)
	default:
		fmt.Println("Buffer empty")
	}

}
