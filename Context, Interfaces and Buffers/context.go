//a sample program to show how to use the context package

package main

import (
	"context"
	"fmt"
	"time"
)

// create a context that does something
func doSomething(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("Doing something")
			time.Sleep(2 * time.Second)
		}
	}
}

func main() {
	//create a context that is cancellable
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	go doSomething(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("Main context is Done")
	}
}
