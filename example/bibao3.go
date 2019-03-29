package main

import "fmt"

//利用闭包
func Fibo3(n int) int {
	if n < 0 {
		return -1
	} else {
		f := Fibonacci()
		result := 0
		for i := 0; i < n; i++ {
			result = f()
		}
		return result
	}
}

func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	fmt.Print(Fibo3(1))
	fmt.Print(Fibo3(2))
	fmt.Print(Fibo3(3))
	fmt.Print(Fibo3(4))
}
