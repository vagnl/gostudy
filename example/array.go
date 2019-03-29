package main

import "fmt"

func main() {
	var arr1 [4]int
	arr1[0] = 0
	arr1[1] = 1
	arr1[2] = 2
	arr1[3] = 3
	fmt.Println(arr1)
	arr1[3] = 100
	fmt.Println(arr1)

	var arr2 []int
	arr2=append(arr2, 1)

	fmt.Println(len(arr2))
	fmt.Println(arr2)

	var arr3=[5]int{1,2,4,5,6}
	arr4 := [...]int{4,5,6,7,8}

	s1 :=[]int{1,2,3}
	fmt.Printf("%T %v \n",s1,s1)
	fmt.Printf("%T %v \n",arr3,arr3)

	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println(arr3==arr4)

	var p *[5]int =&arr3
	fmt.Println(*p)

	//// for 遍历
	//var arr = [5]int{1,2,3,4,5}
	//
	//for i :=0; i < len(arr); i++ {
	//	fmt.Println(arr[i])
	//}
	//
	//// for range 遍历
	//for _, item := range arr {
	//	fmt.Println(item)
	//}

	//modify
	modify1(&arr3)
	modify2(arr4)
	fmt.Println(arr3)
	fmt.Println(arr4)

}

//数组作为参数
func modify1(arr *[5]int){
	arr[0]=100
}


func modify2(arr [5]int){
	arr[0]=100
}

