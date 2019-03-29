package main

import (
	"fmt"
	"time"
)

type count int

func NewCount(i int) count {
	return count(i)
}

//func main() {
//	//c := NewCount(10)
//	//fmt.Println(c)
//
//	var a chan int
//	if a == nil {
//		fmt.Println("channel a is nil, going to define it")
//		a = make(chan int)
//		fmt.Printf("Type of a is %T", a)
//	}
//}

func hello(done chan bool) {
	fmt.Println("Hello world goroutine")
	time.Sleep(4*time.Second)
	done <- true
}
func main() {
	done := make(chan bool)
	go hello(done)
	//<-done
	fmt.Println("main function")
}


