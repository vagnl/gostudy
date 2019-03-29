package main

import (
	"fmt"
	"time"
)

func hello4(done chan bool) {
	fmt.Println("hello4 go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello4 go routine awake and going to write to done")
	done <- true
}
func main() {
	done := make(chan bool)
	fmt.Println("Main going to call hello4 go goroutine")
	go hello4(done)
	//<-done
	fmt.Println("Main received data")
}


