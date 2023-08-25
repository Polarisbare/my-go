package main

import "fmt"

// 数组排序小---->大
func sortArray(arr []int) []int {
	for i := 0; i <= len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// 打印数组长度和容量
func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
func main() {
	var arr = []int{1, 4, 3, 2, 5, 8, 7, 6}
	fmt.Println(arr)
	fmt.Println(sortArray(arr))
	fmt.Println("==========分割线============")
	printSlice(arr)
	// make([]T, len, cap)   make函数创建分片 capacity容量作为可选参数 这里 length 是数组的长度并且也是切片的初始长度
	slice1 := make([]int, 5, 1000)
	slice1 = append(slice1, 1)
	fmt.Println(slice1)
}
