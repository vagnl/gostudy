package main

import (
	"fmt"
)


func main()  {
	var p2 =new(MyStudent)
	p2.Age=34
	p2.Name="bob"
	fmt.Printf("%T %v\n",*p2,*p2)

	stu:=MyStudent{
	Name:"dd",
	Age:45,
	}
	fmt.Println(stu)
}
