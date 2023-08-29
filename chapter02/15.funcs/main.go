package main

import "fmt"

var tall = 1.8
var weight = 70

// 代码作用域    赋值可以改变原值，重新声明，外面不会读取    取值只能从外到内   不能从内到外
func main() {
	tall, weight := 1.8, 70
	fmt.Println(tall, weight)
	//calcBMI()
	add := func(a float64, b float64) float64 {
		return a + b
	}
	result := add(1.0, 4.000)
	fmt.Println(result)
}

func calcBMI() {
	//tall, weight := "1.8", 70
	fmt.Println(tall, weight)
}
