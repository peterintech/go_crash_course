package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var myJson = `{"Name":"Gideo","Age":29}`

	user := User{"Peter", 23}
	user2 := User{"Andrew", 23}

	err := json.Unmarshal([]byte(myJson), &user2)
	if err != nil {
		return
	}
	fmt.Println("Name: ", user2.Name, "Age: ", user2.Age)

	data, err := json.Marshal(user)
	if err != nil {
		return
	}
	fmt.Println(string(data))

	if err != nil {
		return
	}

}
