package main

import (
	"context"
	"fmt"
	"time"
)

// context is used for:
// controlling timeouts & deadlines
// WithTimeout(), WithDeadline()

// cancelling go routines
// cancel()

// passing metadata across your Go app
// WithValue()

func main() {
	exampleTimeout()
	exampleWithValue()
}

func exampleTimeout() {
	ctx := context.Background()

	ctxWtimeout, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	done := make(chan struct{})
	go func() {
		time.Sleep(3 * time.Second)
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("Called the API")
	case <-ctxWtimeout.Done():
		fmt.Println("oh no my timeout expired", ctxWtimeout.Err())
		// do some logic to handle this
	}

}

func exampleWithValue() {
	const UserKey int = 0
	ctx := context.Background()

	ctxWvalue := context.WithValue(ctx, UserKey, "123")

	if userID, ok := ctxWvalue.Value(0).(string); ok {
		fmt.Println("this is the userID: ", userID)
	} else {
		fmt.Println("this is a protected route - no userID found!")
	}
}
