package main

import "fmt"

func main()  {
	s1:=[]int{1,2,3,4}

	s2 := s1[0:2]

	fmt.Println(s2)

	s3:=make([]int,3,6)
	fmt.Println("%T %v",s3,s3)
	s3[0]=1
	s3=append(s3,s1...)
	fmt.Println(cap(s3))

	s4:=s1[2:4]
	fmt.Println(s4)

	s1[3]=100
	fmt.Println(s1)
	fmt.Println(s3)
	fmt.Println(s4)


}
