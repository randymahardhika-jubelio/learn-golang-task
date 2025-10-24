package main

import (
	"fmt"
)

// Task 1: FizzBuzz
func fizzBuzz() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

// Task 2: Slice Summation
func sumSlice(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

func main() {
	// Run FizzBuzz
	fizzBuzz()

	// Example usage of sumSlice
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("Sum of slice:", sumSlice(nums))
}