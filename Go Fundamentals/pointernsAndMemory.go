package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// demonstrate pointers
	var a int = 42
	var p *int = &a // p is a pointer to the memory address of a
	fmt.Println(a)  // prints 42
	fmt.Println(p)  // prints the memory address of a
	fmt.Println(*p) // prints the value of a

	// demonstrate unsafe package
	var b int8 = 42
	fmt.Println(b)                    // prints 42
	pb := (*int8)(unsafe.Pointer(&b)) // pb is a pointer to the memory address of b
	*pb = *pb * 2                     // double the value of b using the pointer pb
	fmt.Println(b)                    // prints 84
}
