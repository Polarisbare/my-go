package main

import "fmt"

func main() {
	//第一种写法
	//for i := 1; i <= 100; i++ {
	//	fmt.Println(i)
	//}
	////第二种
	//j := 1
	//for ; j <= 5; j++ {
	//	fmt.Println(j)
	//}
	////第三种
	//k := 0
	//for k < 5 {
	//	fmt.Println("你好")
	//	k++
	//}
	////第四种
	//a := 0
	//for {
	//	a++
	//	if a <= 3 {
	//		fmt.Println("退出", a)
	//	}
	//}
	m := 0
	for {
		m++
		if m >= 3 {
			fmt.Println(m)
			break
		}
		if m%2 == 0 {
			fmt.Println("被continue了")
			continue
		}
		fmt.Println("结束循环")

	}

}
