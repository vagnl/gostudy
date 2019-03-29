package main

import (
	"fmt"
	"time"
)

func hello3(done chan bool) {
	fmt.Println("hello3 go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello3 go routine awake and going to write to done")
	done <- true
}
func main() {
	done := make(chan bool)
	fmt.Println("Main going to call hello3 go goroutine")
	go hello3(done)
	<-done
	fmt.Println("Main received data")
}
