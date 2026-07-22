package tasks

type Task struct {
	Title string
	Done  bool
	age   int
}

func New(title string, done bool, age int) Task {
	return Task{title, done, age}
}
