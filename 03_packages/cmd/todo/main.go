package main

import (
	"fmt"

	"github.com/peterintech/go_crash_course/03_packages/internal/task"
)

func main() {
	task := tasks.New("Good food", false, 18)
	fmt.Println("task title is: ", task.Title)
	fmt.Println("task status is: ", task.Done)
	fmt.Println("task age is: ", task)
	fmt.Println("first commit")
}
