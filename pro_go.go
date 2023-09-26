package main

import (
	"fmt"
	"time"
)

// func test() {
// 	msg := make(chan string)
// 	fmt.Println("This is test........")
// 	go func() {
// 		msg <- "Test function "
// 		// return <-msg
// 	}()

// 	receiver := <-msg
// 	fmt.Println(receiver)

// }

// func longProcess() string {
// 	return "This longProcess() is completed"
// }

// func one() {
// 	fmt.Println("One")
// }

// func two() {
// 	fmt.Println("Two")
// }

// func three() {
// 	fmt.Println("Three")
// }

// func foo(n int) {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(n, ":", i)
// 		amt := time.Duration(rand.Intn(250))
// 		time.Sleep(time.Millisecond * amt)
// 	}
// }

// func sender(c chan string) {
// 	for i := 0; ; i++ {
// 		c <- "Hello World"
// 		c <- "W3Engineers Ltd."
// 		// channel argument will assign string value
// 	}
// }

// func receiver(c chan string) {
// 	for {
// 		// msg := <-c
// 		// select

// 		select {
// 		case check := <-c:
// 			if check == "Hello World" {
// 				fmt.Println("Hello World")
// 			}
// 			fmt.Println(<-c)
// 		}
// 		// fmt.Println(<-c)
// 		time.Sleep(time.Second * 1)
// 	}
// }

// func sum1(n int) {
// 	sum := 0
// 	for i := 0; i < n; i++ {
// 		sum += n
// 	}
// 	fmt.Println(sum)
// }

// func sum2(5)
// func facto2(n int) int {
// 	facto := 1
// 	for i := 1; i <= n; i++ {
// 		facto = facto * i
// 	}
// 	return facto
// }

// func facto(n int, ch chan<- int) {
// 	//ch chan<- int  only send data to channel
// 	facto := 1
// 	for i := 1; i <= n; i++ {
// 		facto = facto * i
// 	}
// 	ch <- facto
// }

func factorial(n int, ch chan<- int) {
	fact := 1
	for i := 1; i <= n; i++ {
		fact = fact * i
	}
	ch <- fact

}

func factorial2(n int) int {
	fact := 1
	for i := 1; i <= n; i++ {
		fact = fact * i
	}
	return fact
}

func main() {
	var value, k int
	ch := make(chan int)

	timeStart := time.Now()
	for i := 1; i <= 20; i++ {
		go factorial(i, ch)
		value = <-ch
		fmt.Println("Factorial of ", i, "is :", value)
	}

	fmt.Printf("Using Single goroutine took %s\n", time.Since(timeStart))
	fmt.Println("Multiple Goroutine: ")
	timeStart = time.Now()
	for i := 1; i <= 5; i++ {
		go factorial(i, ch)
		value = <-ch
		fmt.Println("Factorial of ", i, "is :", value)
	}

	for i := 6; i <= 10; i++ {
		go factorial(i, ch)
		value = <-ch
		fmt.Println("Factorial of 6 to 10 ", i, "is :", value)
	}

	for i := 11; i <= 15; i++ {
		go factorial(i, ch)
		value = <-ch
		fmt.Println("Factorial of 11 to 15", i, "is :", value)
	}

	for i := 16; i <= 20; i++ {
		go factorial(i, ch)
		value = <-ch
		fmt.Println("Factorial of 16 to 20", i, "is :", value)
	}

	fmt.Printf("Using Multiple goroutine took %s\n", time.Since(timeStart))

	fmt.Println("Without using Goroutine")
	timeStart = time.Now()
	for i := 1; i <= 20; i++ {
		k = factorial2(i)

		fmt.Println("Factorial of ", i, " is\t", k)
	}

	fmt.Printf("Without goroutine took %s\n", time.Since(timeStart))

	// 	var n int
	// 	ch_data := make(chan int)
	// 	fmt.Println("Enter Value for search factorial:")
	// 	fmt.Scan(&n)
	// 	go facto(n, ch_data)
	// 	result := <-ch_data
	// 	fmt.Println("Factorial of\t", n, " is:\t", result)

	// 	fmt.Println("\nFactorial Between number 1 to 20 is: ")

	// 	fmt.Println("Without using Goroutine")
	// 	timeStart := time.Now()
	// 	for j := 1; j <= 20; j++ {
	// 		result = facto2(j)

	// 		fmt.Println("Factorial of ", j, " is\t", result)
	// 	}

	// 	fmt.Printf("Without goroutine took %s\n", time.Since(timeStart))
	// 	timeStart = time.Now()

	// 	for j := 1; j <= 20; j++ {
	// 		go facto(j, ch_data)
	// 		result := <-ch_data
	// 		fmt.Println("Factorial of \t", j, "is \t: ", result)
	// 	}
	// 	fmt.Printf("Using goroutine took %s\n", time.Since(timeStart))

	// =======================================================conf/start=====================================================
	// ch := make(chan int)

	// fmt.Println("Main Go Routine")
	// ch <- 1 // 1
	// ch <- 2 // 2
	// ch <- 6 // 3

	// go func() {
	// 	fmt.Println("First Go Start:")
	// 	ch <- 31 //4
	// 	ch <- 41 //5
	// 	ch <- 51 //6
	// 	ch <- 71 //7
	// 	tenth := <-ch
	// 	fmt.Println(tenth)
	// 	fmt.Println("First Go End")
	// }()

	// go func() {
	// 	fmt.Println("Second Go Start:")
	// 	ch <- 32 //8
	// 	ch <- 42 //9
	// 	ch <- 52 //10
	// 	ch <- 72 //11
	// 	ninth := <-ch
	// 	fmt.Println(ninth)
	// 	fmt.Println("Second Go End")
	// }()

	// go func() {
	// 	fmt.Println("Third Go Start")
	// 	ch <- 33 // assign 11 to fourth fourth = 31, 12
	// 	ch <- 43 //33
	// 	eighth := <-ch
	// 	fmt.Println(eighth)
	// 	fmt.Println("Third GO End")
	// }()

	// first := <-ch
	// second := <-ch
	// third := <-ch
	// fourth := <-ch //33
	// fifth := <-ch
	// sixth := <-ch
	// seventh := <-ch
	// fmt.Println("Before Calling Go:")
	// fmt.Println(first, second, third, fourth, fifth, sixth, seventh)
	// fmt.Println("After Calling Go:")

	// ====================================================================conf-end========================================================
	// go func() {
	// 	ch <- 1
	// 	ch <- 2
	// 	ch <- 3
	// 	ch <- 4
	// }()

	// first := <-ch
	// second := <-ch
	// fmt.Println(first)  //1
	// fmt.Println(second) //2
	// fmt.Println(second) //2
	// fmt.Println(second) //2

	// sum1(5)
	// go sum2(5)
	// time.Sleep(time.Second * 2)
	// fmt.Println(ch)
	// time.Sleep(time.Second * 2)
	// sum := 0
	// slice1 := []int{1, 2, 3, 4, 5}
	// for i := 0; i < len(slice1); i++ {
	// 	sum += slice1[i]
	// }
	// fmt.Println(sum)

	// mychannel := make(chan int, 2)
	// fmt.Println("Value of chan", mychannel)
	// fmt.Printf("Type of chan: %T\n", mychannel)
	// go func() {
	// 	m := make(map[string]int)
	// 	m["one"] = 3000
	// 	val2 := 2
	// 	mychannel <- val2
	// 	mychannel <- m["one"]
	// 	val1 := <-mychannel
	// 	// fmt.Println(<-mychannel)
	// 	fmt.Println(val1)
	// 	// time.Sleep(time.Second * 2)

	// }()
	// time.Sleep(time.Second * 1)
	// fmt.Println(<-mychannel)
	// fmt.Println(<-mychannel)

	// fmt.Println(<-mychannel)

	// defer func() {
	// 	str := recover()
	// 	fmt.Println(str)
	// }()
	// panic("PANIC")
	// panic("PANIC")
	// str := recover()
	// fmt.Println(str)

	// element := map[int]map[string]string{
	// 	1: map[string]string{
	// 		"Symbol": "H",
	// 		"Name":   "Hydrogen",
	// 		"State":  "Gas",
	// 	},
	// 	2: map[string]string{
	// 		"Symbol": "He",
	// 		"Name":   "Hellium",
	// 	},
	// 	3: map[string]string{
	// 		"Symbol": "Li",
	// 		"Name":   "Lithium",
	// 		"State":  "Solid",
	// 	},
	// }

	// // fmt.Println(element[1]["Symbol"])
	// if ele, ok := element[1]; ok {
	// 	fmt.Println(ele["Name"], ele["State"])
	// }

	// elements := map[string]string{
	// 	"H":  "Hydrogen",
	// 	"He": "Hellium",
	// 	"Li": "Lithium",
	// }

	// fmt.Println(elements["H"])

	// elements := map[string]map[string]string{
	// 	"H": map[string]string{
	// 		"name":  "Hydrogen",
	// 		"state": "gas",
	// 	},
	// 	"He": map[string]string{
	// 		"name":  "Helium",
	// 		"state": "gas",
	// 	},
	// 	"Li": map[string]string{
	// 		"name":  "Lithium",
	// 		"state": "solid",
	// 	},
	// }
	// if el, ok := elements["Li"]; ok {
	// 	fmt.Println(el["name"], el["state"])
	// }
	// c1 := make(chan string)
	// c2 := make(chan string)

	// go func() {
	// 	for {
	// 		c1 <- "From 1"
	// 		time.Sleep(time.Second * 2)
	// 	}
	// }()

	// go func() {
	// 	for {
	// 		c2 <- "From 2"
	// 		time.Sleep(time.Second * 1)
	// 	}
	// }()

	// go func() {
	// 	for {
	// 		select {
	// 		case msg1 := <-c1:
	// 			fmt.Println("MSG 1:", msg1)
	// 		case msg2 := <-c2:
	// 			fmt.Println("MSG 2:", msg2)
	// 		case <-time.After(time.Second):
	// 			fmt.Println("Timeout")

	// 		}
	// 	}
	// }()

	// for i := 0; i < 10; i++ {
	// 	go foo(i)
	// }

	// c := make(chan string, 2)
	// go sender(c)   // sending channel as parameter
	// go receiver(c) // sending channel as parameter

	// var input string
	// fmt.Scanln(&input)

	// // slice operation

	// var s []string
	// s = make([]string, 3) //slice
	// c := make([]string, 3)
	// c = []string{"Hello", "World", "GO", "2023"}
	// fmt.Println(c)

	// copy(s, c)
	// s = append(s, "!!")
	// fmt.Println(s)

	// l := c[:2]
	// fmt.Println(l)

	// l = c[1:3]
	// fmt.Println(l)
	// l = c[3:]
	// fmt.Println(l)

	// // Map operation

	// m := make(map[string]int)
	// m["one"] = 1
	// m["two"] = 2
	// m["three"] = 3

	// fmt.Println(m)

	// v1 := m["one"]
	// v2 := m["two"]

	// fmt.Println(v1, v2)

	// channel data passing

	// fmt.Println("Before Calling Main Function From Main")
	// go test()

	// time.Sleep(time.Second * 1)
	// fmt.Println("After calling test function from main")
	// // time.Sleep(time.Second)

	// // go longProcess()

	// go func() {
	// 	send := longProcess()
	// 	ch := make(chan string)
	// 	ch <- send
	// 	receive := <-ch
	// 	fmt.Println(receive)
	// }()

	// defer one()
	// defer two()
	// defer three()

	// ch := make(chan int)
	// msg := make(map[string]int)
	// // var s []string
	// // s = make([]string, 10)

	// msg["one"] = 1
	// msg["two"] = 2
	// msg["three"] = 3

	// ch <- msg["one"]
	// ch <- msg["two"]
	// ch <- msg["three"]

	// rcv := make(map[int]int)

	// rcv[0] = <-ch
	// rcv[0] = <-ch
	// rcv[2] = <-ch

	// fmt.Println(rcv[0], rcv[1], rcv[2])

}
