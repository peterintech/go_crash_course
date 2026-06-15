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
