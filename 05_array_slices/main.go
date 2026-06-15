package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	// Arrays are fixed-size collections of elements of the same type
	var fruitList [4]string

	// Assigning values to the array
	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	fruitList[2] = "Grape"
	fruitList[3] = "Mango"
	fmt.Println(fruitList) // Output: [Apple Orange Grape Mango]

	fmt.Println(fruitList[1]) // Output: Orange

	//Declare and assign
	fruitList2 := [4]string{"Apple", "Orange", "Grape", "Mango"}
	fmt.Println(fruitList2) // Output: [Apple Orange Grape Mango]

	// Slices are dynamic-size collections of elements of the same type (mostly used than arrays)
	fruitSlice := []string{"Apple", "Orange", "Grape", "Mango"}
	fmt.Println(len(fruitSlice)) // Output: 4
	fmt.Println(fruitSlice[1:3]) // Output: [Orange Grape]
}
