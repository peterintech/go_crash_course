package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

// exer 1
func Withdraw(balance, amount float64) (float64, error) {
	if amount < 0 {
		return 0, fmt.Errorf("Error: %v is negetive", amount)
	}
	if amount > balance {
		return 0, fmt.Errorf("Insufficient funds!")
	}
	return balance - amount, nil
}

// exer2
func ReadConfig(filename string) ([]byte, error) {
	open, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("reading file %s: %w", filename, err)
	} else {
		return open, nil
	}
}

// exer 3
func SaveNote(filename, note string) error {
	data := []byte(note)

	err := os.WriteFile(filename, data, 0666)
	if err != nil {
		return fmt.Errorf("Error creating note: %w", err)
	}
	return nil
}

// exer 4
func LoadNote(filename string) (string, error) {
	note, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("Error reading %s: %w", filename, err)
	}
	// don't i have to close the note before returning, i don't know how to do that
	return string(note), nil
}

// exer 5
type Product struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// exer 7
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// exer 8
func Process() {
	//runs first
	fmt.Println("1")

	defer fmt.Println("2") //runs last

	fmt.Println("3") //runs second

	defer fmt.Println("4") //runs 4th

	fmt.Println("5") // runs 3rd
}

// exer11
func ProcessOrder(ctx context.Context) error {
	ctxWtime, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	select {
	case <-time.After(4 * time.Second):
		fmt.Println("Done")
		return nil

	case <-ctxWtime.Done():
		return fmt.Errorf("%w", ctxWtime.Err())
	}
}

// exer12
type Order struct {
	Customer string  `json:"customer"`
	Amount   float64 `json:"amount"`
}

func saveOrder(filename string) error {
	order := Order{}
	order_file, err := ReadConfig(filename)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	json.Unmarshal(order_file, &order)
	if order.Amount < 0 {
		return fmt.Errorf("Insufficient funds!")
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		3*time.Second,
	)
	defer cancel()

	select {
	case <-time.After(2 * time.Second):
		fmt.Printf("\nOrder saved, Amount: %v, Customer: %v\n", order.Amount, order.Customer)
		return nil

	case <-ctx.Done():
		return fmt.Errorf("Saving order failed: %w", ctx.Err())
	}
}

func main() {
	// exer 1: i feel like I'm being repetitive with my error handling here, i feel like have to do else after an if, is this the convention?
	balance, err := Withdraw(1000, 1500)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(balance)

	//exer 2
	file, err := ReadConfig("order.json")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(file)

	//exer 3
	SaveNote("notes.txt", "Learning Go is fun")

	//exer 4
	note, err := LoadNote("notes.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Note says: %q\n", note)

	//exer 5
	prd := Product{
		Name:  "Keyboard",
		Price: 25000,
	}
	res, err := json.Marshal(prd)

	fmt.Println(string(res))

	//exer 6
	data := []byte(`
		{
			"name":"Mouse",
			"price":15000
		}
	`)

	json.Unmarshal(data, &prd)

	fmt.Printf("Product Name: %v\nPrice: %v\n", prd.Name, prd.Price)

	//exer 7
	user := User{}
	user_read, err_read := ReadConfig("user.json")
	if err_read != nil {
		fmt.Println(err)
	}
	json.Unmarshal(user_read, &user)

	fmt.Printf("Name: %v, Email: %v", user.Name, user.Email)

	//exer 9
	defer fmt.Println("A") //runs last
	defer fmt.Println("B") //runs 3rd
	defer fmt.Println("C") //runs 2nd

	fmt.Println("Start") //runs first

	//exer 10
	ctx, cancel := context.WithTimeout(
		context.Background(),
		3*time.Second,
	)
	defer cancel()

	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Done")

	case <-ctx.Done():
		fmt.Println("Timed out1: ", ctx.Err())
	}

	//exer11
	process_err := ProcessOrder(ctx)
	if process_err != nil {
		fmt.Println(process_err)
	}

	//exer12
	err_order := saveOrder("order.json")
	if err_order != nil {
		fmt.Println(err)
	}

	//quiz
	//1. Go wants you to treat errors as values
	//2. for error handling
	//3. file not found is just a string not an error, instead we do errors.New("file not found")
	//4. it wraps the error so we can know what went wrong
	//5. when you want to check if the error is same
	//6. error.is compares the error value, error.as compares the error type
	//7. because you want to perform the close action last
	//8. taking the memory address helps it to update the original user in the memory instead of just creating a copy of user
	//9. it solves the problem of controlling timeouts and deadlines (like abort stuffs) and help to pass metadata across the app
	//10. defer cancel() performs a cleanup over the context, when context is mounted, a lot of stuffs are done under the hood that needs to be cleaned to avoid memory leaks
	// senior thinking challenge
	// just like i did in exercise 12, the json can enter the application either from a file or directly from an api, from a file we first read the file using the os.ReadFile, from an api we take it directly to the app, the next we do is to model the kind of data we are expecting, for this example we'll model the customer, amount json in a struct for example:
	// type Order struct {
	// 	Customer string `json:"customer"`
	// 	Amount   string `json:"amount"`
	// }
	// we initialize this struct: order:=Order{}
	//and then we json.Unmarshal(the data in []byte type, &order)
	// validation will happen in several places which include, the file reading, the json conversion to struct
	// for the context, the parent should have the context initialized via context.Background() or r.Context() provided by the http handler this is then used to pass to the db and everything related to that guy so everything is brought down at once upon failure or timeout

}
