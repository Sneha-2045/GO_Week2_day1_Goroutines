package main

import (
	"fmt"
	"time"
)

func worker(name string) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("%s - Task %d\n", name, i)
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	go worker("Worker 1")
	go worker("Worker 2")
	go worker("Worker 3")

	time.Sleep(3 * time.Second)
}