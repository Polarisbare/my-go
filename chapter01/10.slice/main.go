package main

import "fmt"

// 修改切片内容 []int 不可以转 []byte可以和string胡同
func main() {
	var a string = "hello"
	var b string = "您好"
	fmt.Println(a)
	fmt.Println(b)
	aBytes := []byte(a)  //字母
	aBytes2 := []rune(b) //汉字也可以
	fmt.Println("a", aBytes, len([]byte(a)))
	fmt.Println("b", aBytes2, len([]rune(b)))
	fmt.Println("修改切片内容")
	aBytes[0] = 'H'
	aBytes2[0] = 'K'
	a = string(aBytes)
	b = string(aBytes2)

	fmt.Println(a)
	fmt.Println(b)

}
