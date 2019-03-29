package main

import "fmt"

func main()  {

	s1 :=make([]int,3,10)
	fmt.Println(s1)

	s2 :=make([]int,5)
	fmt.Println(cap(s2))

	m1 := make(map[string]int,5)
	m1["ss"]=1
	m1["rr"]=2
	m1["tt"]=3
	m1["yy"]=4
	m1["uu"]=5
	m1["oo"]=6
	//delete(m1, "s")
	//m1=append(m1,)
	fmt.Println(m1)
	fmt.Println(len(m1))

	m2 := map[string]int{"ee":3,"77":0}
	fmt.Println(m2)


	urls := make(map[string]string,0)
	urls["baidu"] = "www.baidu.com"
	urls["google"] = "www.google.com"
	urls["csdn"] = "www.csdn.net"
	fmt.Println(len(urls))

	v,ok:=urls["baid"]
	fmt.Println(v,ok)

	//names := make([]string, len(urls))
	//for key, _ := range urls {
	//	names = append(names, key)
	//}
	//fmt.Println(names,len(names))



}
