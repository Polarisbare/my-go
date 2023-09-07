package main

import (
	"learn.go/11.practice/01.fatrate.refactor/calc"
	"log"
)

type Calc struct {
	continental string
}

func (Calc) BMI(person *Person) error {
	bmi, err := calc.CalcBMI(person.tall, person.weight)
	if err != nil {
		log.Panicln("bmi计算错误", err)
		return err
	}
	person.bmi = bmi
	return nil
}
func (Calc) FatRate(person *Person) error {
	person.fatRate = calc.CalcFatRate(person.bmi, person.age, person.sex)
	return nil
}
