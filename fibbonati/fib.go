package fibbonati

import (
	"fmt"
	"sync"
)

func fib(num int) int {
	if num <= 0 {
		return 0
	} else if num <= 2 {
		return num
	} else {
		return fib(num-1) + fib(num-2)
	}
}

func calc(start, end int) []int {
	results := []int{}
	for i := start; i < end; i += 1 {
		result := fib(i)
		results = append(results, result)
	}
	return results
}

func calcWithChannel(start, end int, channel chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := start; i < end; i += 1 {
		result := fib(i)
		channel <- result
	}
}

func FibbonatiExample() {
	results := []int{}
	channel := make(chan int)
	var wg sync.WaitGroup
	goroutines := 3
	gap := 10

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		start := i * gap
		end := (i + 1) * gap
		go calcWithChannel(start, end, channel, &wg)
		// go func() {
		// data := calc(start, end)
		// results = append(results, data...)
		// wg.Done()
		// }()
	}

	go func() {
		wg.Wait()
		close(channel)
	}()

	for result := range channel {
		results = append(results, result)
	}

	// wg.Wait()

	fmt.Println(results)
}
