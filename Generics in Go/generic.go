// demonstrate concept of generics in go
// From an array, print the maximum sum of a subarray and also print the subarray, if the array is of string, display the longest string
package main

// parameterized type
type T interface{}

// generic function to check if the array consists int or string, find the maximum sum of a subarray or print the longest string of the given array
func maxSumorLongestString(arr []T) {
	// check if the array consists of int or string
	if _, ok := arr[0].(int); ok {
		// array consists of int, find the maximum sum of a subarray
		maxSum := 0
		for i := 0; i < len(arr); i++ {
			sum := 0
			for j := i; j < len(arr); j++ {
				sum += arr[j].(int)
				if sum > maxSum {
					maxSum = sum
				}
			}
		}
		println("Maximum sum of a subarray is: ", maxSum)

	} else if _, ok := arr[0].(string); ok {
		// array consists of string and print the longest string
		longestString := ""
		for i := 0; i < len(arr); i++ {
			for j := i; j < len(arr); j++ {
				if len(arr[j].(string)) > len(longestString) {
					longestString = arr[j].(string)
				}
			}
		}
		println("Longest string is: ", longestString)
	}
}

func main() {
	// array of integers
	arr1 := []T{1, 2, 3, -2, 5}
	// array of strings
	arr2 := []T{"Hello", "World", "Go", "Programming", "Language"}
	maxSumorLongestString(arr1)
	maxSumorLongestString(arr2)
}
