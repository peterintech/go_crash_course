package main

import (
	"fmt"
	"math"
	"github.com/peterintech/go-crash-course/03_packages/strutil"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(math.Max(5, 10))
	fmt.Println(strutil.Reverse("hello"))
}
