package main

import (
	"flag"
	"fmt"
)

var mode = flag.String("mode", "", "process mode")

func main() {

	// 解析命令行参数
	flag.Parse()
	// 输出命令行参数
	fmt.Println(*mode)

	str := new(string)

	*str = "呵呵大"

	fmt.Println(*str)
}
