package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	convertType()
	fmt.Println("finish")
}
func convertType() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic啦", r)
			debug.PrintStack() //	输出一下调用栈
		}
	}()
	var a any
	a = "string aaa"
	b := a.(int)
	fmt.Println(b)
}
