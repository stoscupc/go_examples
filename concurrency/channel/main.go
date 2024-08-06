package main

import (
	"fmt"
	"sync"
	"time"
)

func addTen(value int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	ch <- value + 10
}

func main() {
	ch := make(chan int, 10)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go addTen(i, ch, &wg)
	}

	wg.Wait()
	close(ch)

	for result := range ch {
		fmt.Println("Received result: ", result)
	}
}
