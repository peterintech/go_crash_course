package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1, num2 int) int {
	return num1 + num2
}

func getSumProduct(num1, num2 int) (int, int) {
	sum := num1 + num2
	product := num1 * num2
	return sum, product
}

func main() {
	// defer is more like: the first shall be the last

	defer fmt.Println("Hi! Peter Here")
	fmt.Println(greeting("Alice"))
	fmt.Println(getSum(5, 10))

	sum, product := getSumProduct(5, 2)

	fmt.Printf("The sum is %d \nThe product is %d\n", sum, product)
}
