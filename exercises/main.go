package main

import (
	"fmt"
)

func multiply(a int, b int) int {
	return a * b
}

func userInfo(name string, age int) (string, int) {
	return name, age
}

func main() {
	//exer 1
	name, age, height := "John", 30, 5.9

	fmt.Printf("Name: %s, Age: %d, Height: %.1f\n", name, age, height)

	//exer 2

	fmt.Printf("Result: %d\n", multiply(5, 3))

	//exer 3
	userName, userAge := userInfo("Alice", 25)
	fmt.Printf("User Info - Name: %s, Age: %d\n", userName, userAge)

	//exer 4
	const pi float32 = 3.14
	fmt.Printf("Value of Pi: %.2f\n", pi)
}
