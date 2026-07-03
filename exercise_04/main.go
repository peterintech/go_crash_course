package main

import (
	"fmt"

	"github.com/peterintech/go_crash_course/exercise_04/mathutil"
	"github.com/peterintech/go_crash_course/exercise_04/user"
)

func main() {
	// exercise 1
	a, b := 10, 6

	sum := mathutil.Add(a, b)
	subtract := mathutil.Subtract(a, b)
	multiply := mathutil.Multiply(a, b)

	fmt.Printf("Arithemetic calculation for %d and %d\n", a, b)
	fmt.Printf("Sum: %d, Difference: %d, Product: %d", sum, subtract, multiply)

	// exercise 2
	// password is lowercase to keep it private so other packages cannot access it
	user1 := user.User{Name: "Peter", Email: "ogbonnapete210@gmail.com"}
	pass := user1.SetPassword("passWorth")
	fmt.Println(pass)

	if user1.CheckPassword("passorth") {
		fmt.Println("The password is correct 👍")
	} else {
		fmt.Println("The password is incorrect 😒")
	}

	// exercise 3
	// a go.mod file is created, it contains the go version, the module name and list of dependencies, its important to track dependencies similar to package.json in js

	//exercise 4
	// I'll choose the structure of project/user and project/product, inside project/user, I'll put the user.go and email.go so there under package user, then inside project/product, I'll put payment.go, invoice.go and product.go so they are under package product, this way I have grouped them based on responsibility
}

// exercise 5
// Product contains the price, invoice and name
type Product struct {
	invoice, price float64
	name           string
}

// CalculateDiscount calculates the discount of the product
func CalculateDiscount(price, percentage float64) float64 {
	return price * (percentage / 100)
}

/*
Part 2
1. A package is a collection of related files who share similar behaviour and responsiblity, these files exist in the same folder and package name convention is the name of the folder they all belong to

2. A package is a collection of related files while a module is a collection of related packages

3. Yes it can, I can have a calculator package and a user package inside one module, i can also have external package, hence the need for a go.mod file where these packages (external ones) are managed

4. Because its generic, its not descriptive of the responsibility, those names will most likely conflict with other packages/ variable/ funcions and also they start with capital letters, its ideal to start package name with small letters

5. Its lowercase no camelCase, It doesn't conflict with other package name, It's descriptive of its responsiblilty

6. If its name starts with a capital letter

7. User, Create and Email are the only ones exported

8. For simplicity and to reduce the number of keywords

9. go.sum is like an integrity checker that checks if packages are tampered with

10. you use a workspace when you're dealing with more than one module (when you have more than onn go.mod file)

Part 3
I'm going to scrap out what i currently have and restructure the codebase like this:
ecommerce/user, ecommerce/orders, ecommerce/payments, ecommerce/auth. Having a total of 4 packages, the user folder & package will have the email and anything relating to the user, the payments folder & package will contain everything related to payments, thesame pattern follows for the auth and orders folder & package
This architecture makes the codebase idiomatic

Part 4
1. the name is too long and it uses Cabab case, lowercase is more accepted, If i'm to rename it'll be just user

2. the name Utilites is generic, i can conflict with other variables/packages/function, also its vague/ ambigous of what the package is about

3. go's convention is to use a letter or two for the customer, we can replace customer ---> c, so its func (c Customer) Save()

4. Its not a good documentation because the comment ought to start with "Add" the name of the function, if i'm to rewrite
// Add gives the sum of two numbers
func Add(a, b int) int

5. To keep things clean and idiomatic

Part 5
1. main.go should not contain business logic, we should hide business logic in packages and manage what we can export like name private stuffs to start with small letter so they are not exportable and just remain there, for example password

2. a good documentation should make that possible he can read via go doc

3. Not everyone might want their age to be known

4. It'll help me know what each package is meant for

explainn the bonus challenge
*/
