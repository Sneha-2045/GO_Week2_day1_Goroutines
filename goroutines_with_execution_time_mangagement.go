package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 5; i++ {
		fmt.Printf("Worker %d -> %d\n", id, i)
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	start := time.Now()

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()

	fmt.Println("Execution Time:", time.Since(start))
}