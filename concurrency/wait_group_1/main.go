package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Create a new WaitGroup
	var wg sync.WaitGroup

	// Increment the WaitGroup counter
	wg.Add(1)

	go func() {
		fmt.Println("Starting...")
		time.Sleep(2 * time.Second)
		fmt.Println("Done!")

		// Decrement the WaitGroup counter
		wg.Done()
	}()

	// Wait until the WaitGroup counter is 0
	wg.Wait()
}
