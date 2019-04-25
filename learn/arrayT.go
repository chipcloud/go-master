package main

import (
	"fmt"
)

func main() {

	fmt.Println("数组练习01")
	var test = [3]int{}

	test[0] = 1
	test[1] = 2
	test[2] = 3

	for i := 0; i < 3; i++ {
		fmt.Println(i, test[i])
	}

	fmt.Println("数组练习02")
	var array = [...]string{"tom", "nike", "adidas"}

	for k, v := range array {
		fmt.Println(k, v)
	}

	fmt.Println("切片练习01")
	var vip = [30]int{}

	for i := 0; i < 30; i++ {
		vip[i] = i + 1
	}

	fmt.Println(vip[20:30])
	fmt.Println(vip[:20])
	fmt.Println(vip[:5])
	fmt.Println(vip[:])

	// 声明字符串切片
	var strList []string
	// 声明整型切片
	var numList []int
	// 声明一个空切片
	var numListEmpty = []int{}
	// 输出3个切片
	fmt.Println(strList, numList, numListEmpty)
	// 输出3个切片大小
	fmt.Println(len(strList), len(numList), len(numListEmpty))
	// 切片判定空的结果
	fmt.Println(strList == nil)
	fmt.Println(numList == nil)
	fmt.Println(numListEmpty == nil)

	a := make([]int, 2)
	b := make([]int, 2, 10)
	fmt.Println(a, b)
	fmt.Println(len(a), len(b))

	var ap []int

	ap = append(ap, 1)

	ap = append(ap, 2, 3)

	//for i:= 0; i< 10; i++ {
	//	ap := append(ap, i)
	//
	//	fmt.Println("len: %d  cap: %d pointer: %p\n", len(ap), cap(ap), ap)
	//}
}
