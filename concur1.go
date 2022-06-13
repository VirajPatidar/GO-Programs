package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	fmt.Println("Hello from main")
	// Increment the WaitGroup counter.
	wg.Add(5)
	go func() {
		fmt.Println("Hello from the other side")
		// Decrements the WaitGroup counter.
		wg.Done()
	}()

	// Waits till the WaitGroup counter reaches 0.
	wg.Wait()
}
