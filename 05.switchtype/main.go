package main

import (
	"fmt"
	"reflect"
)

func main() {
	var money interface{} = 10
	//var money interface{} = 10.0
	//特殊情况下会用到   interface
	//var money interface{} = "10"

	switch newMoney := money.(type) {
	case int:
		tmpMoney := newMoney + 3.0
		fmt.Println("money 是 int", tmpMoney)
		fmt.Println("int")
	case int64:
		fmt.Println("int64")
	case float64:
		fmt.Println("float64")
	case float32:
		fmt.Println("float32")
	default:
		fmt.Println("未知")
	}
	fmt.Println("结束", reflect.TypeOf(money))
	money = 3
	fmt.Println("结束", reflect.TypeOf(money))
}
