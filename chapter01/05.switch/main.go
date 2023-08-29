package main

import (
	"fmt"
)

func main() {
	var money int = 20

	switch money {
	case 20:
		fmt.Println("点个外卖")
		fallthrough //写状态机（ 状态流转时候会用 ）
	case 200:
		fmt.Println("下馆子")
	case 2000:
		fmt.Println("吃米其林")
	default:
		fmt.Println("在想想")
	}
}
