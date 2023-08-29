package main

import "fmt"

func main() {
	chinese, english, math := getScoresOfStudent()
	fmt.Println(chinese, english, math)
}

// 函数名字    +    括号     +   （代表返回的值）
func getScoresOfStudent() (chinese int, english int, math int) {
	chinese, english, math = 90, 88, 99
	return
}
