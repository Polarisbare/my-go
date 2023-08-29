package main

import (
	"fmt"
)

// 根据身高和体重计算出当前同学的状况

func main() {
	var name string
	fmt.Print("姓名：")
	fmt.Scanln(&name)
	var weight float64
	//Print直接输出
	fmt.Print("体重（千克）：")
	fmt.Scanln(&weight)
	var tall float64
	fmt.Print("身高（米）：")
	fmt.Scanln(&tall)
	var bmi float64 = weight / (tall * tall)
	var age int
	fmt.Print("年龄（岁）：")
	fmt.Scanln(&age)
	var sexWeight int
	var sex string
	fmt.Print("性别（男/女）：")
	fmt.Scanln(&sex)
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
				fmt.Println("您目前是： 偏瘦")
			} else if fatRate > 0.1 && fatRate <= 0.16 {
				fmt.Println(name, "您目前是： 标准")
			} else if fatRate > 0.16 && fatRate <= 0.21 {
				fmt.Println(name, "您目前是： 一点胖")
			} else if fatRate > 0.21 && fatRate <= 0.26 {
				fmt.Println(name, "您目前是： 肥胖")
			} else {
				fmt.Println(name, "您目前是：小猪啦")
			}
		} else if age >= 40 && age <= 59 {
			if fatRate <= 0.11 {
				fmt.Println(name, "您目前是： 偏瘦")
			} else if fatRate > 0.11 && fatRate <= 0.17 {
				fmt.Println(name, "您目前是： 标准")
			} else if fatRate > 0.17 && fatRate <= 0.22 {
				fmt.Println(name, "您目前是： 一点胖")
			} else if fatRate > 0.22 && fatRate <= 0.27 {
				fmt.Println(name, "您目前是： 肥胖")
			} else {
				fmt.Println(name, "您目前是：小猪啦")
			}
		} else {
			fmt.Println("男性年龄太大了无法判断")
		}
	} else {
		if age < 18 {
			fmt.Println("女孩：未成年人无法判断")
		} else if age >= 18 && age <= 39 {
			if fatRate <= 0.2 {
				fmt.Println(name, "您目前是： 偏瘦")
			} else if fatRate > 0.2 && fatRate <= 0.27 {
				fmt.Println(name, "您目前是： 标准")
			} else if fatRate > 0.27 && fatRate <= 0.34 {
				fmt.Println(name, "您目前是： 一点胖")
			} else if fatRate > 0.34 && fatRate <= 0.39 {
				fmt.Println(name, "您目前是： 肥胖")
			} else {
				fmt.Println(name, "您目前是：小猪啦")
			}
		} else if age >= 40 && age <= 59 {
			if fatRate <= 0.21 {
				fmt.Println(name, "您目前是： 偏瘦")
			} else if fatRate > 0.21 && fatRate <= 0.28 {
				fmt.Println(name, "您目前是： 标准")
			} else if fatRate > 0.28 && fatRate <= 0.35 {
				fmt.Println(name, "您目前是： 一点胖")
			} else if fatRate > 0.35 && fatRate <= 0.40 {
				fmt.Println(name, "您目前是： 肥胖")
			} else {
				fmt.Println(name, "您目前是：小猪啦")
			}
		} else {
			fmt.Println("女性年龄太大了无法判断")
		}
	}
	//	判断用户是否继续测
	var whetherContinue string
	fmt.Print("是否录入下一个（y/n）？")
	fmt.Scanln(&whetherContinue)
	if whetherContinue == "y" {
		var name string
		fmt.Print("姓名：")
		fmt.Scanln(&name)
		var weight float64
		//Print直接输出
		fmt.Print("体重（千克）：")
		fmt.Scanln(&weight)
		var tall float64
		fmt.Print("身高（米）：")
		fmt.Scanln(&tall)
		var bmi float64 = weight / (tall * tall)
		var age int
		fmt.Print("年龄（岁）：")
		fmt.Scanln(&age)
		var sexWeight int
		var sex string
		fmt.Print("性别（男/女）：")
		fmt.Scanln(&sex)
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
					fmt.Println("您目前是： 偏瘦")
				} else if fatRate > 0.1 && fatRate <= 0.16 {
					fmt.Println(name, "您目前是： 标准")
				} else if fatRate > 0.16 && fatRate <= 0.21 {
					fmt.Println(name, "您目前是： 一点胖")
				} else if fatRate > 0.21 && fatRate <= 0.26 {
					fmt.Println(name, "您目前是： 肥胖")
				} else {
					fmt.Println(name, "您目前是：小猪啦")
				}
			} else if age >= 40 && age <= 59 {
				if fatRate <= 0.11 {
					fmt.Println(name, "您目前是： 偏瘦")
				} else if fatRate > 0.11 && fatRate <= 0.17 {
					fmt.Println(name, "您目前是： 标准")
				} else if fatRate > 0.17 && fatRate <= 0.22 {
					fmt.Println(name, "您目前是： 一点胖")
				} else if fatRate > 0.22 && fatRate <= 0.27 {
					fmt.Println(name, "您目前是： 肥胖")
				} else {
					fmt.Println(name, "您目前是：小猪啦")
				}
			} else {
				fmt.Println("男性年龄太大了无法判断")
			}
		} else {
			if age < 18 {
				fmt.Println("女孩：未成年人无法判断")
			} else if age >= 18 && age <= 39 {
				if fatRate <= 0.2 {
					fmt.Println(name, "您目前是： 偏瘦")
				} else if fatRate > 0.2 && fatRate <= 0.27 {
					fmt.Println(name, "您目前是： 标准")
				} else if fatRate > 0.27 && fatRate <= 0.34 {
					fmt.Println(name, "您目前是： 一点胖")
				} else if fatRate > 0.34 && fatRate <= 0.39 {
					fmt.Println(name, "您目前是： 肥胖")
				} else {
					fmt.Println(name, "您目前是：小猪啦")
				}
			} else if age >= 40 && age <= 59 {
				if fatRate <= 0.21 {
					fmt.Println(name, "您目前是： 偏瘦")
				} else if fatRate > 0.21 && fatRate <= 0.28 {
					fmt.Println(name, "您目前是： 标准")
				} else if fatRate > 0.28 && fatRate <= 0.35 {
					fmt.Println(name, "您目前是： 一点胖")
				} else if fatRate > 0.35 && fatRate <= 0.40 {
					fmt.Println(name, "您目前是： 肥胖")
				} else {
					fmt.Println(name, "您目前是：小猪啦")
				}
			} else {
				fmt.Println("女性年龄太大了无法判断")
			}
		}
	} else {
		return
	}
}
