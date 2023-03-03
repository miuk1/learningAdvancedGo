package main

import (
	"fmt"
)

func main() {
	// if statement
	num := 10
	if num > 0 {
		fmt.Println("The number is positive")
	} else if num < 0 {
		fmt.Println("The number is negative")
	} else {
		fmt.Println("The number is zero")
	}

	// switch statement
	grade := "B"
	switch grade {
	case "A":
		fmt.Println("Excellent")
	case "B":
		fmt.Println("Good")
	case "C":
		fmt.Println("Fair")
	default:
		fmt.Println("Poor")
	}

	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// while loop
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// infinite loop with break statement
	k := 0
	for {
		if k == 5 {
			break
		}
		fmt.Println(k)
		k++
	}

	// continue statement
	for l := 0; l < 5; l++ {
		if l == 3 {
			continue
		}
		fmt.Println(l)
	}

	// nested loop with labeled break statement
outer:
	for m := 0; m < 5; m++ {
		for n := 0; n < 5; n++ {
			if m == 3 && n == 3 {
				break outer
			}
			fmt.Printf("(%d, %d)\n", m, n)
		}
	}

	// defer statement
	defer fmt.Println("World")
	fmt.Println("Hello")

}
