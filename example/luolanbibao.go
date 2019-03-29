package main


import "fmt"

func f()(ret int)  {
	defer func() {
		ret++
	}()
	return 10
}

func main() {
	fmt.Println(f())



}
