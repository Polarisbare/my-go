package main

import "fmt"

// 递归很浪费性能 尽量要用循环来代替
func main() {
	//	斐波那契数列计算核心
	a := fibonacci(1)
	fmt.Println(a)
	fmt.Println("------------------")
	guess(0, 100)
}

func fibonacci(n uint) uint {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// 二分法  猜数字程序
func guess(left, right uint) {
	guessed := (left + right) / 2
	var getFromInput string
	fmt.Println("我猜的是：", guessed)
	fmt.Println("如果高了，输入1，如果低了，输入0； 猜对了输入6")
	fmt.Scanln((&getFromInput))
	switch getFromInput {
	case "1":
		if left == right {
			fmt.Println("你是不是变更数字了，耍赖不是好行为哦！")
			break
		}
		guess(left, guessed-1)
	case "0":
		if left == right {
			fmt.Println("你是不是变更数字了，耍赖不是好行为哦！")
			break
		}
		guess(guessed+1, right)
	case "6":
		fmt.Println("你猜的是：", guessed)
	}

}
