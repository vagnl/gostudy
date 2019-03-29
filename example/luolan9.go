package main

import "fmt"

type dosomething interface {
	work() int
	study() int
}

type Student struct {
	Name string
	Age  int
}

func (stu Student) work() int {
	fmt.Printf("name:%s,age:%d -- work\n", stu.name, stu.age)
	return 3
}

func (stu Student) study() {
	fmt.Printf("name:%s,age:%d -- study\n", stu.name, stu.age)
}

func main() {
	s1 := new(Student)
	s1.Name = "zhang san"
	s1.Age = 23
	s2 := new(Student)
	s2.Name = "Li si"
	s2.Age = 23

	i := s1.work // 作为变量使用
	a := i()
	fmt.Println(a)
	s2.study()
}
