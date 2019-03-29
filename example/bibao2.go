package main

import "fmt"

// fibonacci 函数会返回一个返回 int 的函数。
func fibonacci() func() int {
    var x int =0
    var y int =1
    var temp int
	res:=func()int{
		temp=x+y
		x=y
		y=temp
		return x
	}
	return res
}



func main() {
	k:=0
	fmt.Println(k)
	f := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

