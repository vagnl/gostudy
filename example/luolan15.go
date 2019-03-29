package main

import (
	"fmt"
)

func main() {
	fn := func(a int, b int, z int64) int {
		return a * b
	}

	fmt.Println(fn(4, 4, 5)) //打印fn的类型, 即func(int, int, float64) bool
}
