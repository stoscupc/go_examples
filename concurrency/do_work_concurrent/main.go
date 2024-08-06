package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func delay(d int) {
	time.Sleep(time.Duration(d) * time.Second)
}

func do_work(d int) {
	fmt.Println("Starting work...")
	delay(d)
	fmt.Println("Work done! Job took", d, "seconds")
	wg.Done()
}

var jobs = []int{1, 2, 3, 4}

func main() {
	start := time.Now()

	for _, job := range jobs {
		wg.Add(1)
		go do_work(job)
	}

	wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s\n", elapsed)
}
