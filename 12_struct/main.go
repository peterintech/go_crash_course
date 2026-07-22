package main

import (
	"fmt"
	"strconv"
)

// define a person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// shorter way
type Person2 struct {
	firstName, lastName, city, gender string
	age                               int
}

// two reason why we use pointer receiver
// The first is so that the method can modify the value that its receiver points to.
// The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct,

// This is a Go convention:
// If one method needs a pointer receiver
// All methods should use pointer receivers

// greeting method
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried method (pointer receiver)
func (p *Person) getMarried(newLastName string) {
	if p.gender == "Male" {
		return
	} else {
		p.lastName = newLastName
	}
}

func main() {
	// init person using struct
	p1 := Person{firstName: "Bob", lastName: "Smith", city: "New York", gender: "Male", age: 25}

	p2 := Person{"Samantha", "Smith", "Boston", "Female", 5}
	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p2.greet())
	p2.hasBirthday()
	p2.getMarried("Johnson")
	fmt.Println(p2.greet())
}
