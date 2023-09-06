package main

import "fmt"

func main() {
	a, b := 1, 2
	//	&指针         *反转指针取到真实值
	add(&a, &b)
	fmt.Println(a)

	c := &a //	c的类型是*int，c 指向a的盒子，*c就可以拿到a里面的东西
	d := &c //	d的类型是**int，d指向c的盒子，d 本身是指针，它存的东西也是指针
	fmt.Println("d=", d, "*d=", *d, "**d=", **d)

	f1 := add
	f1(&a, &b)
	fmt.Println("f1, add =", a)
	f1p1 := &f1 //	f1p1 = *func(int,int)
	(*f1p1)(&a, &b)
	fmt.Println("f1p1, add=", a)
}
func add(a, b *int) {
	*a = *a + *b
}
