package main

import "fmt"

const isCool bool = true

func main() {
	var name = "John Doe"
	var age int32 = 30

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)

	fmt.Printf("%T\n", age)
	fmt.Printf("%t\n", isCool)

	//shot hand declaration
	// this can only be used inside a function
	email := "peter@gmail.com"
	name = "Peter Parker"

	name, email = "Bruce Wayne", "bruce@gmail.com"
	
	fmt.Println("Name:", name)
	fmt.Println("Email:", email)
}