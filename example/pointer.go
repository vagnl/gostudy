package main

import "fmt"

func main(){


	i, j := 42, 37

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	fmt.Println(*p+1)


	m := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("m ==", m)
	m = append(m, 2, 3, 4)
	for i := 0; i < len(m); i++ {
		fmt.Printf("m[%d] == %d\n", i, m[i])
	}

}
