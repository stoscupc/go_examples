package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func fetchStatus(url string, ch chan<- string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	ch <- fmt.Sprintf("%s - %s", url, resp.Status)
	wg.Done()
}

func delay(d int) {
	time.Sleep(time.Duration(d) * time.Second)
}

func main() {
	start := time.Now()
	bufferSize := 10

	for i := 1; i <= 10; i++ {
		ch := make(chan string, bufferSize)
		fmt.Println("\nBatch: ", i)
		for i := 0; i < bufferSize; i++ {
			wg.Add(1)
			go fetchStatus("http://github.com", ch)
		}

		wg.Wait()
		delay(1)
		close(ch)

		for status := range ch {
			fmt.Println(status)
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("\nExecution time: %s\n", elapsed)
}
