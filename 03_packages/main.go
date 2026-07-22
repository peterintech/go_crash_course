package main

import (
	"fmt"
	"math"

	// "rsc.io/quote"
	tasks "github.com/peterintech/go_crash_course/03_packages/internal/task"
	"github.com/peterintech/go_crash_course/03_packages/strutil"
)

func main() {
	todo := tasks.New("task", false, 20)
	fmt.Println("Hello, World!", todo.Title)
	fmt.Println(math.Max(5, 10))
	// fmt.Println(quote.Hello())
	fmt.Println(strutil.Reverse("hello"))
}
