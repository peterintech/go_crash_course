package main

import "fmt"

func fizzBuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	// // long method
	// i := 1
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// //short method
	// for j := 1; j <= 10; j++ {
	// 	fmt.Println(j)
	// }

	// // fizzbuz
	// fizzBuzz()
	for i := 0; i <= 7; i++ {
		fmt.Println("masgi")
	}
}
