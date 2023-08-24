package main

import (
	"fmt"
)

func main() {
	// 数组 数组长度固定 1-数组长度不能变 2-类型固定 3-动态赋值
	var age int = 35
	fmt.Println(age)
	var ages [5]int = [5]int{35, 32, 33, 36, 37}
	fmt.Println(ages)
	var ages2 [5]int = [5]int{35, 32, 33, 36, 37}
	fmt.Println(ages2)
	ages3 := [...]int{35, 32, 33, 36, 37}
	fmt.Println(ages3)
	ages4 := [3]int{}
	ages4[0] = 1111
	ages4[1] = 222
	ages4[2] = 444
	fmt.Println("ages4", ages4)
	// 常规循环
	for i := 0; i < len(ages4); i++ {
		fmt.Println(ages4[i])
	}
	// 循环方法 range 写法一
	for i := range ages4 {
		fmt.Println(ages4[i])
	}
	// 循环方法 range 写法二
	for i, val := range ages4 {
		//fmt.Println(ages4[i], "===>", i, "-->", val)
		fmt.Printf("%d\t%d\n", i, val)
	}

}
