package main

import "testing"

func TestCase1Part1(t *testing.T) {
	inputRecord("王强", 0.38)
	inputRecord("王强", 0.32)
	{
		randOfWQ, fatRateOfWQ := getRand("王强")
		if randOfWQ != 1 {
			//	把案例标记成失败但是还是继续会往下跑
			//t.Fail()
			//	结束之后直接退出
			t.Fatalf("预期王强第一，但是得到的是：%d", randOfWQ)
			return
		}
		if fatRateOfWQ != 0.32 {
			t.Fatalf("预期王强的体脂是0.32，但的得到的是 %f", fatRateOfWQ)
		}
	}
}

func TestCase1(t *testing.T) {
	inputRecord("王强", 0.38)
	inputRecord("王强", 0.32)
	{
		randOfWQ, fatRateOfWQ := getRand("王强")
		if randOfWQ != 1 {
			//	把案例标记成失败但是还是继续会往下跑
			//t.Fail()
			//	结束之后直接退出
			t.Fatalf("预期王强第一，但是得到的是：%d", randOfWQ)
			return
		}
		if fatRateOfWQ != 0.32 {
			t.Fatalf("预期王强的体脂是0.32，但的得到的是 %f", fatRateOfWQ)
		}
	}
	inputRecord("李静", 0.28)
	{
		randOfLJ, fatRateOfLJ := getRand("李静")
		if randOfLJ != 1 {
			//	把案例标记成失败但是还是继续会往下跑
			//t.Fail()
			//	结束之后直接退出
			t.Fatalf("预期李静第一，但是得到的是：%d", randOfLJ)
			return
		}
		if fatRateOfLJ != 0.28 {
			t.Fatalf("预期王强的体脂是0.28，但的得到的是 %f", fatRateOfLJ)
		}
	}

}

func TestCase2(t *testing.T) {

}
