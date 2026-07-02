package main

import "fmt"

/*1. Create a Rectange type with methods Area() and Perimeter()
	Expected outcome: Rectangle type with methods Area()
	Area() returns width × height.
2. Create an interface Printer with Print() string
	then implement it for two different structs.
3. Create two receivers for the same type,
	one value receiver and one pointer receiver, and explain why mutation works in one but not the other
*/

// exercise 1
type Rectange struct {
	width, height float64
}

func (r Rectange) Area() float64 {
	return r.height * r.width
}
func (r Rectange) Perimeter() float64 {
	return 2 * (r.height + r.width)
}

// exercise 2
type Printer interface {
	Print() string
}

type Book struct {
	name  string
	pages int
}
type Invoice struct {
	name   string
	amount float64
}

func (b Book) Print() string {
	return fmt.Sprintf("The name of the book is %v, it has %v pages", b.name, b.pages)
}

func (b Invoice) Print() string {
	return fmt.Sprintf("Hi %v, your invoice is ₦%v", b.name, b.amount)
}

func PrintInfo(p Printer) string {
	return p.Print()
}

// exercise 3
type Counter struct {
	Value int
}

func (c Counter) IncrementValue() {
	c.Value++
}
func (c *Counter) IncrementPointer() {
	c.Value++
}

// exer 4
type Notifier interface {
	Notify(message string) error
}

type EmailNotifier struct {
	email string
}
type SmsNotifier struct {
	number string
}

func (e EmailNotifier) Notify(message string) error {
	if message != "" {
		return fmt.Errorf("Hi, %v, you said this: %s", e.email, message)
	}

	return fmt.Errorf("Couldn't process, Please define a message")

}
func (s SmsNotifier) Notify(message string) error {
	if message != "" {
		return fmt.Errorf("You have a new message: %s, to your number: %v", message, s.number)
	}
	return fmt.Errorf("Couldn't process, Please define a message")
}

func Broadcast(notifiers []Notifier, message string) error {
	for _, notifier := range notifiers {
		return notifier.Notify(message)
	}
	return fmt.Errorf("Couldn't process, Please define a message")
}

func main() {
	// exer 1
	r1 := Rectange{9, 5}
	fmt.Println("Area of rectange: ", r1.Area())
	fmt.Println("Perimeter of rectange: ", r1.Perimeter())

	// exer 2
	book := Book{"Pete's diary", 512}
	invoice := Invoice{"Tobe", 300000}

	fmt.Println(PrintInfo(invoice))
	fmt.Println(PrintInfo(book))

	//exer 3
	count := Counter{5}
	fmt.Println("Count before increments: ", count.Value)

	count.IncrementValue()
	fmt.Println("Count after IncrementValue: ", count.Value)
	count.IncrementPointer()
	fmt.Println("Count after IncrementPointer: ", count.Value)

	// the incrementPointer changes the value because via the pointer it gets the memory address to the count struct and modifies it, while the incrementValue doesn't change the value because it doesn't have access to the address of count to modify it, every mutation there is just scoped to its function

	//exer 4 - not sure on this one (give me a detailed review on the correction of this question)
	email := EmailNotifier{"ogbonnapete210@gmail.com"}
	sms := SmsNotifier{"09012345678"}

	arr := []Notifier{email, sms}

	fmt.Println(Broadcast(arr, "Afa where my money?"))
}
