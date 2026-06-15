package main

import "fmt"

func main() {
	fmt.Println("Maps in Go")

	//Define a map
	email := make(map[string]string)

	email2 := map[string]string{
		"Bob":    "bob@gmail.com",
		"Alice":  "alice@gmail.com",
		"Sharon": "sharon@gmail.com",
	}

	//Assign key values
	email["Bob"] = "bob@gmail.com"
	email["Alice"] = "alice@gmail.com"

	//Access values
	fmt.Println(email["Bob"])
	fmt.Println(email["Alice"])
	fmt.Println(email2)

	//Delete from map
	delete(email, "Bob")
	fmt.Println(email)
}
