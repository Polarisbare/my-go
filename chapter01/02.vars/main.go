package main

import (
	"fmt"
	"math"
)

func main() {
	var hello string = "hello,golang!"
	var world = "world"
	fmt.Println(hello, world)
	小数 := 1.233

	fmt.Println(小数)

	var int3, int4 uint = 33, 44
	fmt.Println(int3, int4)

	var ho, ver float64 = 3, 4.88999
	var sc = ho * ver
	fmt.Println(ho * ver)
	fmt.Println(sc)

	var int6 uint = math.MaxUint64
	fmt.Println(int6)
	var int7 int = int(int6)
	fmt.Println(int7)

	var (
		int5, int9 = 1, 2
	)
	fmt.Println(int5, int9)
}
