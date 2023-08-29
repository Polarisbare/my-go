package main

import (
	"fmt"
	"reflect"
)

func main() {
	arr := make([]string, 3, 100)
	fmt.Printf("长度: %d\n", len(arr))
	fmt.Printf("容量: %d\n", cap(arr))
	fmt.Println("default", arr[0])
	fmt.Println("default", arr[1])
	fmt.Println("default", arr[2])

	//	new指针
	i := new(int)
	fmt.Println(reflect.TypeOf(i))

	//	copy 赋值切片
	arr3 := make([]string, 3, 4)
	copy(arr3, arr)
	fmt.Println(arr3)
}
