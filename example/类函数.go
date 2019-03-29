package main

import "fmt"

type Cat1 struct {
	Name    string
	Age     int32
	Address string
}

func (cat *Cat1) Grow1() {
	cat.Age++
}
func main() {
	myCat := Cat1{"Little C", 2, "In the house"}
	myCat.Grow1()
	fmt.Printf("%v", myCat)
}
