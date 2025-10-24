package main

import (
	"fmt"
)

// Task 1: Day of the Week
func dayOfWeek(day int) string {
	switch day {
	case 1:
		return "Monday"
	case 2:
		return "Tuesday"
	case 3:
		return "Wednesday"
	case 4:
		return "Thursday"
	case 5:
		return "Friday"
	case 6:
		return "Saturday"
	case 7:
		return "Sunday"
	default:
		return "Invalid day"
	}
}

// Task 2: Type Inspector
func typeInspector(val interface{}) {
	switch v := val.(type) {
	case int:
		fmt.Println("The type is int")
	case string:
		fmt.Println("The type is string")
	case bool:
		fmt.Println("The type is bool")
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}

func main() {
	// Test Task 1
	fmt.Println(dayOfWeek(1)) // Monday
	fmt.Println(dayOfWeek(7)) // Sunday
	fmt.Println(dayOfWeek(9)) // Invalid day

	// Test Task 2
	typeInspector(42)
	typeInspector("hello")
	typeInspector(true)
	typeInspector(3.14)
}