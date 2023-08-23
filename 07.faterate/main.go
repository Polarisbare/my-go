package main

import (
	"fmt"
)

// 根据身高和体重计算出当前同学的状况

func main() {
	fmt.Println()
	var weight float64 = 65.0
	var tall float64 = 1.82
	var bmi float64 = weight / (tall * tall)
	var age int = 26
	var sexWeight int
	var sex string = "男"
	if sex == "男" {
		sexWeight = 1
	} else {
		sexWeight = 0
	}
	var fatRate float64 = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
	fmt.Println("您的体脂率是", fatRate)
	if sex == "男" {
		if age < 18 {
			fmt.Println("男孩：未成年人无法判断")
		} else if age >= 18 && age <= 39 {
			if fatRate <= 0.1 {
				fmt.Println("目前是： 偏瘦")
			} else if fatRate > 0.1 && fatRate <= 0.16 {
				fmt.Println("目前是： 标准")
			} else if fatRate > 0.16 && fatRate <= 0.21 {
				fmt.Println("目前是： 一点胖")
			} else if fatRate > 0.21 && fatRate <= 0.26 {
				fmt.Println("目前是： 肥胖")
			} else {
				fmt.Println("目前是：小猪啦")
			}
		} else if age >= 40 && age <= 59 {
			if fatRate <= 0.11 {
				fmt.Println("目前是： 偏瘦")
			} else if fatRate > 0.11 && fatRate <= 0.17 {
				fmt.Println("目前是： 标准")
			} else if fatRate > 0.17 && fatRate <= 0.22 {
				fmt.Println("目前是： 一点胖")
			} else if fatRate > 0.22 && fatRate <= 0.27 {
				fmt.Println("目前是： 肥胖")
			} else {
				fmt.Println("目前是：小猪啦")
			}
		} else {
			fmt.Println("男性年龄太大了无法判断")
		}
	} else {
		// todo 编写女性的体脂率与体脂状态表
		if age < 18 {
			fmt.Println("女孩：未成年人无法判断")
		} else if age >= 18 && age <= 39 {
			if fatRate <= 0.2 {
				fmt.Println("目前是： 偏瘦")
			} else if fatRate > 0.2 && fatRate <= 0.27 {
				fmt.Println("目前是： 标准")
			} else if fatRate > 0.27 && fatRate <= 0.34 {
				fmt.Println("目前是： 一点胖")
			} else if fatRate > 0.34 && fatRate <= 0.39 {
				fmt.Println("目前是： 肥胖")
			} else {
				fmt.Println("目前是：小猪啦")
			}
		} else if age >= 40 && age <= 59 {
			if fatRate <= 0.21 {
				fmt.Println("目前是： 偏瘦")
			} else if fatRate > 0.21 && fatRate <= 0.28 {
				fmt.Println("目前是： 标准")
			} else if fatRate > 0.28 && fatRate <= 0.35 {
				fmt.Println("目前是： 一点胖")
			} else if fatRate > 0.35 && fatRate <= 0.40 {
				fmt.Println("目前是： 肥胖")
			} else {
				fmt.Println("目前是：小猪啦")
			}
		} else {
			fmt.Println("女性年龄太大了无法判断")
		}
	}

}
