package main

import "fmt"

// panic  and  recover
func panicAndRecover() {
	//defer coverPanicPuls() //  可以抓住错误
	//defer coverPanic() //  抓不住
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("代码报错啦")
		}
	}()
	//错误函数
	//panic("出现错误")
	var nameScore map[string]int = nil
	nameScore["无敌"] = 100
}

func coverPanic() {
	func() {
		if r := recover(); r != nil {
			fmt.Println("代码报错啦")
		}
	}()
}
func coverPanicPuls() {

	if r := recover(); r != nil {
		fmt.Println("代码报错啦")
	}

}
