package main

import (
    "fmt"
    "sync"
)

func worker(id int, ch chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    for job := range ch {
        fmt.Printf("Worker %d received job %d\n", id, job)
    }
}

func main() {
    jobs := make(chan int)
    var wg sync.WaitGroup

    // Start 3 worker goroutines
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go worker(i, jobs, &wg)
    }

    // Send jobs to workers
    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs) // Close channel to signal no more jobs

    wg.Wait()
}

