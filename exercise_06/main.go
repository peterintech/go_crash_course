package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// exer 1
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Division by Zero: %v/%v", a, b)
	}
	return a / b, nil
}

// exer 2
func LoadFile(filename string) ([]byte, error) {
	result, err := os.ReadFile(filename)

	if err != nil {
		return nil, fmt.Errorf("reading file %v: %w", filename, err)
	}
	return result, nil

}

// exer3
func CreateFile(filename string, data []byte) error {
	err := os.WriteFile(filename, data, 0666)

	if err != nil {
		return fmt.Errorf("An error: %w", err)
	}
	return nil
}

// exer4
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// exer 6
func Process() {
	//runs first
	fmt.Println("Start")
	//runs last
	defer fmt.Println("Cleanup")
	//runs second
	fmt.Println("Working")
}

// exer7
// I dont know

//Part 2 Quiz
//1. because in go errors are values

//2. it means if and error exists show it stop the program
//3 file not found is just a string, err.Error() is and error, that's why we use the errors.Is instead
//4. to wrap errors
//5. when you want to check if the error is same
//6. error.is compares the error value, error.as compares the error type
//7. you gotta cleanup/close the file you opened and that's where defer saves the dayc
//8. it got to change the original value of the user
//9. It solves the problem of me being able to interrupt an action due to situation, for example someone queries our db and he shuts his system, i should be able to stop that process
//10. for cleanup

//part 3
//1. I'd decode the request json by using json.unMarshal()

func main() {
	var a, b float64 = 10, 7
	result, err := Divide(a, b)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("%v divided by %v is %f \n", a, b, result)

	//exer 2
	file, err := LoadFile("config.go")
	if err != nil {
		fmt.Println(err)
		// return
	}
	fmt.Printf("the file, %v\n", file)

	//exer3
	cerr := CreateFile("notes.txt", []byte("Learning Go is fun!"))
	if cerr != nil {
		fmt.Println(cerr)
	}
	res, err := LoadFile("notes.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("The result: %v\n", res)

	//exer4
	user := User{
		Name: "Peter",
		Age:  25,
	}
	jsonv, err := json.Marshal(user)

	fmt.Printf("The json is: %v\n", jsonv)

	//exer5
	data := []byte(`{
    "name":"Alice",
    "age":30
	}`)
	json.Unmarshal(data, &user)
	fmt.Printf("Name: %v\nAge: %v", user.Name, user.Age)
}
