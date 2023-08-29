package main

import "fmt"

func main() {
	//	map
	/*  map **************注意事项***************
	一：Map 不用实力话就可以读取、删除
	二：Map 必须实例化才可以写入      否则会空指针 崩溃
	三：重复删除相同的 key 不会引起异常

	var m1 map[string]int = nil        nil 表示空  js中的null
	*/
	names := []string{"小强", "周易", "旭东"}
	fr := []float64{28, 18, 19}

	names = append(names, "无敌")
	fr = append(fr, 10)
	for i, name := range names {
		if name == "周易" {
			fmt.Printf("%s,的体脂率是%f\n", name, fr[i])
		}
	}
	fmt.Println("============定义map===========")
	var m1 = map[string]int{}
	m2 := map[string]int{}
	m3 := map[string]int{"王强": 60, "小李": 85, "校长": 100}
	fmt.Println(m1, m2, m3)
	//  利用map查询
	fmt.Println("王强的分数：", m3["王强"])
	//	添加
	m3["无敌"] = 77
	fmt.Println(m3["无敌"])
	//	删除
	delete(m3, "小李")
	fmt.Println(m3["小李"])
	//	修改
	m3["无敌"] = 1000
	fmt.Println(m3["无敌"])

	//	Map的遍历
	for k, v := range m3 {
		fmt.Printf("%s\t%d\n", k, v) //   等价   fmt.Println(k, "=", v)

	}
}
