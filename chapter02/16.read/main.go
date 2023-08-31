package main

import (
	"fmt"
	"github.com/spf13/cobra" // 类型转换的包
	learn_go_tools "learn.go.tools"
	"learn.go/11.practice/01.fatrate.refactor/calc"
)

// 命令行定义
func main() {
	fmt.Println(learn_go_tools.Max(3, 5))
	//	录入
	var (
		name   string
		sex    string
		tall   float64
		weight float64
		age    int
	)
	//	os.Args 读取参数
	cmd := &cobra.Command{
		Use:   "healthcheck",
		Short: "体脂计算器，根据吧啦吧啦计算体脂，给出建议（此处写简短描述）",
		Long:  "这个写详细描述",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("name  ", name)
			fmt.Println("sex   ", sex)
			fmt.Println("tall  ", tall)
			fmt.Println("weight", weight)
			fmt.Println("age   ", age)
			//	计算
			bmi := calc.CalcBMI(tall, weight)
			fatRate := calc.CalcFatRate(bmi, age, sex)
			fmt.Println("fatRate =", fatRate)
			//	评估结果
		},
	}
	cmd.Flags().StringVar(&name, "name", "", "姓名")
	cmd.Flags().StringVar(&sex, "sex", "", "性别")
	cmd.Flags().Float64Var(&tall, "tall", 0, "身高")
	cmd.Flags().Float64Var(&weight, "weight", 0, "体重")
	cmd.Flags().IntVar(&age, "age", 0, "年龄")

	cmd.Execute()
}
