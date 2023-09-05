package calc

import (
	"testing"
)

// 单元测试
func TestCalcBMI(t *testing.T) {
	inputHeight, inputWeight := 1.0, 1.0
	expecteDuput := 1.0
	t.Logf("开始计算，输入：heigh他:%f,weight:%f,期望结果：%f", inputHeight, inputWeight, expecteDuput)
	actualOutput, err := CalcBMI(inputHeight, inputWeight)
	t.Logf("实际得到的：%f,error:%v", actualOutput, err)
	if err != nil {
		t.Fatalf("expecting no err,but got:%v", err)
	}
	if expecteDuput != actualOutput {
		t.Errorf("expecting %f,but got %f", expecteDuput, actualOutput)
	}
}
