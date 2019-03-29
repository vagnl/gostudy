package main


import "fmt"

func Sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // 将和送入 c
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go Sum(a[:len(a)/2], c)
	go Sum(a[len(a)/2:], c)
	x:=<-c
	y:=<-c
	//x, y := <-c, <-c // 从 c 中获取

	fmt.Println(x, y)
}

