package main

import (
	"fmt"
)

// Function example
func sum(a, b int) int {
	return a + b
}

// Struct with method example
type Person struct {
	Name string
	Age  int
}

func (p Person) isAdult() bool {
	return p.Age >= 18
}

func main() {
	// If/else statement example
	x := 10
	if x < 5 {
		fmt.Println("x is less than 5")
	} else if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is equal to 5")
	}

	// For loop example
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Switch statement example
	y := 5
	switch y {
	case 1:
		fmt.Println("y is 1")
	case 5:
		fmt.Println("y is 5")
	default:
		fmt.Println("y is neither 1 nor 5")
	}

	// Function call example
	z := sum(2, 3)
	fmt.Println(z)

	// Struct and method call example
	p1 := Person{Name: "Alice", Age: 25}
	p2 := Person{Name: "Bob", Age: 17}
	fmt.Println(p1.isAdult()) // true
	fmt.Println(p2.isAdult()) // false
}
