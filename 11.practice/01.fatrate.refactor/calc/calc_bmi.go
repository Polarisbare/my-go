package calc

func CalcBMI(tall float64, weight float64) (bmi float64, err error) {
	//	计算bmi
	if tall <= 0 {
		panic("身高不能是0或者负数")
	}
	return weight / (tall * tall), nil
}
