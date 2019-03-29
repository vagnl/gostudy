package main

type MyStudent struct{
	Name string
	Age int
}

func (stu *MyStudent)Set(name string,age int){
	stu.Name = name
	stu.Age = age
}

func (stu *MyStudent)GetAge()(int){
	return stu.Age
}


/*func main(){
	var s MyStudent
	s.Set("tome",23)
	fmt.Println(s)
	fmt.Println(s.GetAge())
}*/
