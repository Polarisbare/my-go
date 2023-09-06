package main

import (
	"fmt"
	"learn.go/11.practice/01.fatrate.refactor/calc"
)

func main() {
	//	go 语言支持面向对象编程，虽然go语言不是严格的面向对象编程语言
	//	go 语言汇总的结构体与其他语言类型中的类在面向对象编程汇总处于同等地位
	//	go语言中没有严格的继承，方法重载，构造函数等
	//	go语言通过接口完成面向对象编程
	//	组装做到很好，紧密联系的组装到一起：须具备 1、高内聚 2、低耦合
	persons := []Person{
		{
			"小强",
			"男",
			1.7,
			70,
			35,
		},
	}
	for _, item := range persons {
		bmi, err := calc.CalcBMI(item.tall, item.weight)
		fmt.Println(bmi, err)
	}
}

type Person struct {
	name   string
	sex    string
	tall   float64
	weight float64
	age    int
}
