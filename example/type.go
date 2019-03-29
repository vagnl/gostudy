package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"time"
)

//bool
//
//string
//
//int  int8  int16  int32  int64
//uint uint8 uint16 uint32 uint64 uintptr
//
//byte // uint8 的别名
//
//rune // int32 的别名// 代表一个Unicode码
//
//float32 float64
//complex64 complex128



var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
	z1 complex64 =3+9i
	i1 int =1
	i2 int8 =2
	i3 int16 =3
	i4 int32=4
	i5 int64=5
	i6 rune=6
	i7 uint=7
	i8 uint8=8
	i9 uint16=9
	i10 uint32=10
	i11 uint64=11
	i12 uintptr=12

	f1 float32 =32.0
	f2 float64 = 55.66


	//var a[4]int ={ 1,2,4,6 }


)

func main() {
	const f = "%T: %v\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
	fmt.Printf(f, z1, z1)

	fmt.Printf(f, i1, i1)
	fmt.Printf(f, i2, i2)
	fmt.Printf(f, i3, i3)
	fmt.Printf(f, i4, i4)
	fmt.Printf(f, i5, i5)
	fmt.Printf(f, i6, i6)
	fmt.Printf(f, i7, i7)
	fmt.Printf(f, i8, i8)
	fmt.Printf(f, i9, i9)
	fmt.Printf(f, i10, i10)
	fmt.Printf(f, i11, i11)
	fmt.Printf(f, i12, i12)

	var i int
	var t float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, t, b, s)


	var x, y int = 3, 4
	var fx  = math.Sqrt(float64(x*x + y*y))
	var z int = int(fx)
	fmt.Println(x, y, z)

	ti := time.Now()
	switch {
	case ti.Hour() < 12:
		fmt.Println("Good morning!")
	case ti.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	defer fmt.Println("world")

	fmt.Println("hello")

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

}

