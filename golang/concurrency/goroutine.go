package main

import (
	"fmt"
	"sync"
)

func sum(arr []int, wg *sync.WaitGroup, ch chan int) {
	total := 0
	for _, val := range arr {
		total += val
	}
	ch <- total
	wg.Done()
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	ch := make(chan int, 2) //A buffered channel of integers is created with a capacity of 2. It will be used to send the sums from each goroutine back to the main function.

	var wg sync.WaitGroup // declared to help wait for the two goroutines to finish their work before the program continues.
	wg.Add(2)

	go sum(arr[:len(arr)/2], &wg, ch)
	go sum(arr[len(arr)/2:], &wg, ch)

	wg.Wait()
	close(ch)

	totalSum := 0
	for val := range ch {
		totalSum += val
	}

	fmt.Println("Total Sum: ", totalSum)
}
