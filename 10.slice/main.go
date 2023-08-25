package main

import "fmt"

// 修改切片内容 []int 不可以转 []byte可以和string胡同
func main() {
	var a string = "hello"
	fmt.Println(a)
	aBytes := []byte(a)
	fmt.Println(aBytes)
	fmt.Println("修改切片内容")
	aBytes[0] = 'H'
	a = string(aBytes)
	fmt.Println(a)

}
