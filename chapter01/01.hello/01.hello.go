package main

import (
	"fmt"
)

func main() {
	var name string = "小强"
	fmt.Println(name)
	var age int = 35
	fmt.Println(age)
	var prices float64 = 100.90
	fmt.Println(prices)
	name = "小强变量测试"
	//age = "hahahhah" 类型不同会报错
}
