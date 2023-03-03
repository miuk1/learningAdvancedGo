package main

import "fmt"

func main() {
	// create an array
	arr := [5]int{1, 2, 3, 4, 5}

	// iterate over the array using a for loop
	fmt.Println("Iterating over an array using a for loop:")
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// create a slice from the array
	slice := arr[1:4]

	// append a value to the slice
	slice = append(slice, 6)

	// iterate over the slice using a for loop
	fmt.Println("Iterating over a slice using a for loop:")
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// iterate over the slice using a range loop
	fmt.Println("Iterating over a slice using a range loop:")
	for index, value := range slice {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// create a struct
	type person struct {
		name string
		age  int
	}

	// create a variable of type person
	p := person{name: "John", age: 30}

	// print the value of the variable
	fmt.Println("Printing the value of a struct:")
	fmt.Println(p)

	// create a pointer to the variable
	pp := &p

	// update the value of the variable using the pointer
	pp.age = 31

	// print the updated value of the variable
	fmt.Println("Printing the updated value of a struct:")
	fmt.Println(p)
}
