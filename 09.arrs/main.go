package main

import "fmt"

// 多维数组
func main() {
	fmt.Println("v1")
	// 难以维护
	xqInfo := [3]string{"小强", "男", "在职"}
	xlInfo := [3]string{"小李", "男", "在职"}
	xsInfo := [3]string{"小苏", "男", "在职"}
	fmt.Println(xqInfo)
	fmt.Println(xlInfo)
	fmt.Println(xsInfo)
	fmt.Println("v2")
	// [3][3] 代表 3行3列
	// 数组长度管理
	newPersonInfo := [3][3]string{
		[3]string{"小强", "男", "在职"},
		[3]string{"小李", "男", "在职"},
		[3]string{"小苏", "男", "在职"},
	}
	for _, val := range newPersonInfo {
		fmt.Println(val)
	}
	fmt.Println("v3")
	// ...支持动态添加
	newPersonInfo2 := [...][3]string{
		[3]string{"小强", "男", "在职"},
		[3]string{"小李", "男", "在职"},
		[3]string{"小苏", "男", "在职"},
		[3]string{"小哈", "男", "在职"},
	}
	for _, val := range newPersonInfo2 {
		fmt.Println(val)
	}
	fmt.Println("v4降维方式")
	// 用降维方式
	for d1, d1val := range newPersonInfo2 {
		for d2, d2val := range d1val {
			fmt.Println(d1, d1val, d2, "val", d2val)
		}
	}
}
