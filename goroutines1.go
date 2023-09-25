package main
import (
"fmt"
"time"
)
func concurrentFunc(label string) {
fmt.Println(label)
time.Sleep(1 * time.Second)
}
func main() {
go concurrentFunc("First process")
go concurrentFunc("Second process")
go concurrentFunc("Third process")
time.Sleep(1 * time.Second) // to wait main function
}
