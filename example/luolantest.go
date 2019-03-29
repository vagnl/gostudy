package main

import "fmt"

func main (){

	const (
		a = iota //a=0
		b = "B"
		c = iota //c=2
		d, e, f = iota, iota, iota //d=3,e=3,f=3
		g = iota //g = 4
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
}
