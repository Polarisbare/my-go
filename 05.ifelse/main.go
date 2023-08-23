package main

import "fmt"

func main() {
	var rmb int
	if rmb <= 20 {
		fmt.Println("点个外卖")
	} else if rmb <= 200 {
		fmt.Println("下馆子")
	} else if rmb <= 2000 {
		fmt.Println("米其林")
	} else {
		fmt.Println("容我想想")
	}

}
