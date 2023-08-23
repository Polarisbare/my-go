package main

import "fmt"

func main() {
	//a, b := 100, 31
	//fmt.Println(a ^ b)
	//fmt.Println(b ^ a)
	//
	//arr := []int{4, 5, 4, 5, 6, 7, 6, 3, 3}
	//result := -1
	//for _, item := range arr {
	//	if result < 0 {
	//		result = item
	//	} else {
	//		result = result ^ item
	//	}
	//}
	//fmt.Println(result)

	a := 2
	b := 10101010

	//a = a + b
	//b = a - b
	//a = a - b
	fmt.Printf("%b", a)
	fmt.Println("b = ", *&b)

}
