package main

import "fmt"

func main() {

	//	结构体的属性在同一个包内均可见
	//	只有公有的结构体、成员变量、成员函数在包外可见
	//	结构体的成员函数执行时，只能通过结构体指针的成员函数进行更改

	var left, right int = 1, 2
	//var op string = "+"
	c := Calculator{
		left:  left,
		right: right,
		//op:    op,
	}
	fmt.Println(c.Add())
	fmt.Println("c.result=", c.result)
}
