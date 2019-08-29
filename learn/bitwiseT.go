package main

import "fmt"

func main() {

	n := 2
	for ; n > 0; n-- {

		if n&1 == 1 {
			fmt.Println(n, "是个奇数")
		} else {
			fmt.Println(n, "是个偶数")
		}

	}

	x := 10

	y := 12

	x = x ^ y // （1）
	y = x ^ y // （2）
	x = x ^ y // （3）

	fmt.Println("x", x)
	fmt.Println("y", y)

	array := [...]int{1, 2, 3, 1, 2, 3, 4}

	i := len(array)
	tmp := array[0]

	for ; i > 0; i-- {
		tmp = tmp ^ array[i-1]
	}

	fmt.Println("唯一的数字是：", tmp)

}
