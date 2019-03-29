package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func main(){
	fmt.Println(time.Now().Unix())
	rand.Seed(4)
	fmt.Println("My favorite number is", rand.Intn(15))

	//RRfmt.Println("rand",rand.Intn(10))

	//fmt.Println(math.Pi)
	//var i  =64.98
	//var i1  = int(i)
	//fmt.Println(i1)

	//g := 0.867 + 0.5i
	//f := 0.143
	//fmt.Println(f)
	//
	//fmt.Println(g)
	//fmt.Println(1i)


	//fmt.Println(Big)

	defer fmt.Println("world")

	fmt.Println("hello")
	fmt.Println("counting")

	//延迟的函数调用被压入一个栈中。当函数返回时， 会按照后进先出的顺序调用被延迟的函数调用。

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

}