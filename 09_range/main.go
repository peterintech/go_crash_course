package main

import "fmt"

func main() {

	// Range is used to iterate over elements in a variety of data structures such as arrays, slices, maps, and channels.
	ids := []int{33, 44, 55, 66, 77}

	//loop through the ids and print them
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	//add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Printf("Sum: %d\n", sum)

	//range with maps
	email := map[string]string{
		"Bob":    "bob@gmail.com",
		"Alice":  "alice@gmail.com",
		"Sharon": "sharon@gmail.com",
	}
	for name, email := range email {
		fmt.Printf("Name: %s, Email: %s\n", name, email)
	}

	for k := range email {
		fmt.Printf("Name: %s\n", k)
	}
}
