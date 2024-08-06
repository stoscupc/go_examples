package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		fmt.Println("About to save int in channel...")
		ch <- 1
		fmt.Println("Saved int in channel...")
	}()

	fmt.Println("About to read int from channel...")
	n := <-ch
	fmt.Println("Read int from channel...", n)
}
