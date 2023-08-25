package main

import "fmt"

// 增加删除
func main() {
	//  range 循环  ****  append 追加  *****
	//  切片：可以动态增加、删除切片中的元素
	a := []int{1, 2}

	fmt.Println("追加元素到a中，a是切片")
	fmt.Println("========================")
	a = append(a[1:], 333)
	fmt.Println(a)
	//	数组
	b := [2]int{1, 2}
	fmt.Println(b)
	xqInfo := []string{"小强", "小路", "小张"}
	for i, val := range xqInfo {
		fmt.Println(i, val)
	}
	fmt.Println("========================")
	fmt.Println("删除切片中的元素")
	a = []int{111, 222, 333, 444, 555}
	//a = append(a[:1], a[2:]...)
	//a = a[:3]
	a = append(a, a[:1]...)
	fmt.Println(a)

}
