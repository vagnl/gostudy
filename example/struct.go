package main

import "fmt"

type Human struct {
	Id   int
	Sex  string
	Name string
}

type Worker struct {
	Salary int
	Base   Human
}

//func NewHuman(id int, sex string, name string) *Human {
//	return &Human{Id: id, Sex: sex, Name: name}
//}
//
//func NewHuman()*Human{
//	return &Human{
//		Id:10,
//		Sex:"uuu",
//		Name:"ddd",
//	}
//}

func (hu *Human) Changename() {
	hu.Name = "change"
}
func main() {
	student1 := Human{
		Id:   23,
		Sex:  "female",
		Name: "anny",
	}

	var student2 Human
	student2.Name = "tny"
	student2.Id = 98
	student2.Sex = "male"

	worker1 := Worker{
		Salary: 6000,
		Base: Human{
			Id:   677,
			Sex:  "female",
			Name: "rose",
		},
	}

	student1.Changename()

	fmt.Printf("%+v\n", student1)
	fmt.Printf("%+v\n", student2)
	fmt.Printf("%+v\n", worker1)

}
