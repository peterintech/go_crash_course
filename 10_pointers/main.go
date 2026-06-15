package main

import "fmt"

func main() {
	a := 10
	b := &a           // b is a pointer to a
	fmt.Println(a, b) // Output: 10 0xc000018060

	// use * to read val from address
	fmt.Println(*b) // Output: 10

	// change value of a using pointer b
	*b = 20
	fmt.Println(a) // Output: 20
}
