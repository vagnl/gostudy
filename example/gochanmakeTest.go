package main

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum //将和送入c
	fmt.Println("sum=",sum)
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}

	c := make(chan int)

	t:=len(a)/2
	fmt.Println(t)

	fmt.Println(a[len(a)/2:])
	fmt.Println(a[:len(a)/2])

	go sum(a[len(a)/2:], c)
	go sum(a[:len(a)/2], c)


	//x, y := <-c, <-c //从c中获取

	//fmt.Println(x, "+", y, "=", x+y)

	fmt.Println("end")
}
