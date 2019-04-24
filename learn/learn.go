package main

import (
	"fmt"
	"math"
)

func main() {

	var name = "王凯"

	var age = 25

	var price = []float32{10.0, 12.0}

	var flag func() bool

	var es struct {
		x int
	}

	var attack = 40
	var defence = 20
	var damageRate float32 = 0.17
	var damage = float32(attack-defence) * damageRate
	fmt.Println(damage)

	hp := 100

	println(name, age, price, flag, es.x)

	fmt.Println(hp)

	//conn, e := net.Dial("tpc", "127.0.0.1:8080")
	//
	//fmt.Print(conn)
	//
	//fmt.Print(e)

	var a = 100

	var b = 120

	a, b = b, a

	fmt.Println("啊", a, b)

	i, _ := GetData()
	_, j := GetData()
	fmt.Println(i, j)

	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)

	const str = `
第一行
第二行
第三行
	`
	fmt.Println(str)
	var by byte = 'a'
	fmt.Printf("%d %T\n", by, by)
	var ni rune = '你'
	fmt.Printf("%d %T\n", ni, ni)

	// 输出各数值范围
	fmt.Println("int8 range:", math.MinInt8, math.MaxInt8)
	fmt.Println("int16 range:", math.MinInt16, math.MaxInt16)
	fmt.Println("int32 range:", math.MinInt32, math.MaxInt32)
	fmt.Println("int64 range:", math.MinInt64, math.MaxInt64)
	// 初始化一个32位整型值
	var abc int32 = 1047483647
	// 输出变量的十六进制形式和十进制值
	fmt.Printf("int32: 0x%x %d\n", abc, abc)
	// 将a变量数值转换为十六进制, 发生数值截断
	cba := int16(abc)
	// 输出变量的十六进制形式和十进制值
	fmt.Printf("int16: 0x%x %d\n", cba, cba)
	// 将常量保存为float32类型
	var c float32 = math.Pi
	// 转换为int类型, 浮点发生精度丢失
	fmt.Println(int(c))
}

type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func GetData() (int, int) {
	return 100, 200
}
