package main

import "fmt"

func main() {
	var cat int = 1
	var str string = "banana"
	fmt.Printf("%p %p", &cat, &str)

	var house = "Malibu Point 10880, 90265"

	ptr := &house

	// 打印ptr的类型
	fmt.Printf("ptr type: %T\n", ptr)
	// 打印ptr的指针地址
	fmt.Printf("address: %p\n", ptr)
	// 对指针进行取值操作
	value := *ptr
	// 取值后的类型
	fmt.Printf("value type: %T\n", value)
	// 指针取值后就是指向变量的值
	fmt.Printf("value: %s\n", value)

	a, b := 1, 2

	swap(&a, &b)

	fmt.Println(a, b)
}

func swap(a, b *int) {

	// 取a指针的值, 赋给临时变量t
	t := *a
	// 取b指针的值, 赋给a指针指向的变量
	*a = *b
	// 将a指针的值赋给b指针指向的变量
	*b = t

}
