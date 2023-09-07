package main

import (
	"learn.go/11.practice/01.fatrate.refactor/calc"
	"log"
)

type Person struct {
	name   string
	sex    string
	tall   float64
	weight float64
	age    int

	bmi     float64
	fatRate float64
}

func (p *Person) calcBmi() error {
	bmi, err := calc.CalcBMI(p.tall, p.weight)
	if err != nil {
		log.Printf("BMi for Person[%s]:%v", p.name, err)
		return err
	}
	p.bmi = bmi
	return nil
}
func (p *Person) calcFatRate() {
	p.fatRate = calc.CalcFatRate(p.bmi, p.age, p.sex)
}
