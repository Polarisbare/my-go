package main

import (
	"fmt"
	"os"

	_ "github.com/spf13/cobra" // 类型转换
)

// 命令行定义
func main() {
	//	录入
	var (
		name   string
		sex    string
		tall   string // float64
		weight string //float64
		age    string //int
	)
	//	os.Args 读取参数
	arguments := os.Args
	name = arguments[1]
	sex = arguments[2]
	tall = arguments[3]
	weight = arguments[4]
	age = arguments[5]
	fmt.Println(arguments)
	fmt.Println("name  ", name)
	fmt.Println("sex   ", sex)
	fmt.Println("tall  ", tall)
	fmt.Println("weight", weight)
	fmt.Println("age   ", age)
	//	计算
	//	评估结果
}
