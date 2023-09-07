package main

import (
	"fmt"
)

func main() {
	var left, right int
	var op string
	fmt.Scanln(&left)
	fmt.Print("")
	fmt.Scanln(&op)
	fmt.Scanln(&right)
	fmt.Print("")
	switch op {
	case "+":
		fmt.Println(left + right)
	case "-":
		fmt.Println(left - right)
	case "*":
		fmt.Println(left * right)
	case "/":
		fmt.Println(left / right)
	case "%":
		fmt.Println(left % right)
	default:
		fmt.Println("不合法")
	}
}

type Calculator struct {
	left, right int
	op          string
}

//Add()
//Sub()
//Multiple()
//Divide()
//Reminder()
