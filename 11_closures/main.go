package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	sum := adder()
	fmt.Println(sum(5))  // Output: 5
	fmt.Println(sum(10)) // Output: 15

	for i := 0; i <= 10; i++ {
		fmt.Println(sum(i))
	}
}
