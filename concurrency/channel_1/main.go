package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 1
	}()

	go func() {
		ch <- 2
	}()

	v1, v2 := <-ch, <-ch

	fmt.Println(v1, v2)
}
