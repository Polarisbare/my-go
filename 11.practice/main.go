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

func main() {
	var arr = []int{1, 4, 3, 2, 5, 8, 7, 6}
	fmt.Println(arr)
	fmt.Println(sortArray(arr))
}
