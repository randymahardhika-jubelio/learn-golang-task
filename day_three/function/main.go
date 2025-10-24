package main

import (
	"fmt"
)

// Task 1: Multiple Return Values
func calculate(a, b int) (int, int) {
	sum := a + b
	product := a * b
	return sum, product
}

// Task 2: Variadic Sum Function
func sumAll(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func main() {
	// Test Task 1
	sum, product := calculate(3, 4)
	fmt.Printf("Sum: %d, Product: %d\n", sum, product)

	// Test Task 2
	fmt.Println("sumAll(1, 2):", sumAll(1, 2))
	fmt.Println("sumAll(5, 10, 15, 20):", sumAll(5, 10, 15, 20))
}