package main

import "fmt"

func (c *Calculator) Add() int {
	tempResult := c.left + c.right
	c.result = tempResult
	fmt.Println("调用了")
	return tempResult
}
func (c Calculator) Sub() int {
	return c.left - c.right
}
func (c Calculator) Multiple() int {
	return c.left * c.right
}
func (c Calculator) Divide() int {
	return c.left / c.right
}
func (c Calculator) Reminder() int {
	return c.left % c.right
}
