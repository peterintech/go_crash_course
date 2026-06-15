package main

import (
	"fmt"
	"math"

	// "rsc.io/quote"
	"github.com/peterintech/go_crash_course/03_packages/strutil"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(math.Max(5, 10))
	// fmt.Println(quote.Hello())
	fmt.Println(strutil.Reverse("hello"))
}
