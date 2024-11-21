package main

import (
	"fmt"
	"sync"
)

func printNumber(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(num)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go printNumber(i, &wg)
	}
	wg.Wait()
}
