package calc

func CalcFatRate(bmi float64, age int, sex string) (fatRate float64) {
	//	计算体脂率
	sexWeight := 0
	if sex == "男" {
		sexWeight = 1
	} else {
		sexWeight = 0
	}
	fatRate = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
	return fatRate
}
