package main

import "fmt"

func main() {
	// Define some variables of different data types
	var a int = 10
	var b float64 = 3.1415
	var c bool = true
	var d string = "Hello, World!"

	// Perform some operations on the variables
	sum := a + int(b)
	diff := int(b) - a
	product := int(b) * a
	quotient := int(b) / a
	remainder := int(b) % a
	greater := a > int(b)
	less := a < int(b)
	equal := a == int(b)
	notEqual := a != int(b)

	// Print the results of the operations
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("d:", d)
	fmt.Println("sum:", sum)
	fmt.Println("diff:", diff)
	fmt.Println("product:", product)
	fmt.Println("quotient:", quotient)
	fmt.Println("remainder:", remainder)
	fmt.Println("greater:", greater)
	fmt.Println("less:", less)
	fmt.Println("equal:", equal)
	fmt.Println("notEqual:", notEqual)
}
