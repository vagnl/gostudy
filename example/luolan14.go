package main

import "fmt"

type Describer interface {
	Describe()
}
type Person1 struct {
	name string
	age  int
}

func (p Person1) Describe() {
	fmt.Printf("%s is %d years old", p.name, p.age)
}

func findType1(i interface{}) {
	switch v := i.(type) {
	case Describer:
		v.Describe()
	default:
		fmt.Printf("unknown type\n")
	}
}

func main() {
	findType1("Naveen")
	p := Person1{
		name: "Naveen R",
		age:  25,
	}
	findType1(p)
}
