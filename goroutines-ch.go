package main

import (
	"fmt"
	"time"
)

// func main() {
// 	messages := make(chan string)

// 	go work(messages)

// 	msg := <-messages
// 	fmt.Println("Channel running on", msg)
// }

// func work(messages chan<- string) {
// 	messages <- "golang"
// }

// func main() {
// 	isDone := make(chan bool, 1)

// 	go work(isDone)

// 	<-isDone
// 	fmt.Println("finished")
// }

// func work(a chan bool) {
// 	fmt.Println("Working...")
// 	time.Sleep(time.Second)
// 	a <- true
// }

func main() {
	c1 := make(chan string, 1)

	//Running longRunningProcess in it's own goroutine and pass
	//back it's response int c1 channel

	go func() {
		text := longRunningProcess()
		c1 <- text

	}()

	//Listen on our channel AND a timeout channel - which
	// ever happens first

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(7 * time.Second):
		fmt.Println("Out of time")
	}
}

func longRunningProcess() string {
	time.Sleep(5 * time.Second)
	return "This Code is completed"
}
