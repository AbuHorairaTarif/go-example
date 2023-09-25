package main

import (
	"fmt"
	"sync"
	"time"
)

// func main() {
// 	fmt.Println("Start Script")

// 	for i := 0; i < 10; i++ {
// 		go printItem(i)
// 	}

// 	fmt.Println("End Script")

// 	// Wait, giving time for the go routines
// 	// to finish.
// 	time.Sleep(10000)
// }

// func printItem(i int) {
// 	fmt.Printf("Print Item: %v\n", (i + 1))
// }

func main() {
	total := 3

	var wg sync.WaitGroup
	wg.Add(total)

	for i := 1; i <= total; i++ {
		go longRunningProcesss(i, &wg)
	}
	wg.Wait()
	fmt.Println("Finished")

}

func longRunningProcesss(sleep int, wg *sync.WaitGroup) {

	// Call Done() using defer as its easiest way to guarantee
	// its called at every exit
	defer wg.Done()

	time.Sleep(time.Duration(sleep) * time.Second)
	fmt.Println("Sleeping for", sleep, "seconds")

}
