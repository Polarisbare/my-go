package main

import (
	"fmt"
)

func main() {
	name := "1111"
	age := 333
	fmt.Printf("%p/n", &age)
	//fmt.Println(name, age)
	age, time := 32, "时间"
	fmt.Printf("%p/n", &age)
	fmt.Println(name, age, time)
}
