package main

import (
	"fmt"
	"strings"
)

// Lesson Two — Control flow, slices, maps, and structs exercise
// 1. Write a function
//  Expected outcome:
//  "go go rust" returns countWords(text string) map[string]int .map[string]int{"go": 2, "rust": 1} .

// 2. Write a function
// topStudent(scores map[string]int) string .
// Expected outcome:
// return the name with the highest score.

// 3. Write a small program that filters even integers from a slice and returns a new slice.
// Expected outcome:
// input [1,2,3,4,5,6] gives [2,4,6]

// 4. Extend the float64 and print all orders above a chosen threshold.

// exer 1
func countWords(text string) map[string]int {
	words := strings.Split(text, " ")
	result := make(map[string]int)

	for _, word := range words {
		result[word]++
	}
	return result
}

// exer 2
func topStudent(scores map[string]int) string {
	highestScore := 0
	highestName := ""
	for name, score := range scores {
		if score > highestScore {
			highestScore = score
			highestName = name
		}
	}
	return highestName
}

// exer 3
func filterEvenIntegers(x []int) []int {
	b := []int{}
	for _, v := range x {
		if v%2 == 0 {
			b = append(b, v)
		}
	}
	return b
}

// exer4
type Order struct {
	ID       int
	Customer string
	Items    []string
	Tags     map[string]bool
	Amount   float64 //added
}

func main() {
	fmt.Println("Exercises..")
	text := "go go rust"

	//exer 1:
	fmt.Printf("\nThe count: %v", countWords(text))

	//exer 2:
	studentScores := map[string]int{
		"peter": 90,
		"alice": 75,
		"uche":  89,
	}
	highestStud := topStudent(studentScores)

	fmt.Printf("\nThe higest Student is %v\n", highestStud)

	// exer 3:
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	even := filterEvenIntegers(a)
	fmt.Printf("The Even version of %v is %v", a, even)

	//exer4
	orders := []Order{
		{
			ID:       1,
			Customer: "Ada",
			Items:    []string{"keyboard", "mouse"},
			Tags:     map[string]bool{"paid": true, "priority": true},
			Amount:   30.06,
		},
		{
			ID:       2,
			Customer: "Linus",
			Items:    []string{"monitor"},
			Tags:     map[string]bool{"paid": false},
			Amount:   20.50,
		},
	}
	for _, order := range orders {
		if order.Amount > 50 {
			fmt.Printf("\nCustomer with the name %v has his amount above 50", order.Customer)
		} else {
			fmt.Printf("\nCustomer with the name %v has his amount below 50", order.Customer)
		}
	}

}
