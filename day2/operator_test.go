package main

import "fmt"

func main() {
	var a int
	a = 10 //修改的是啥？8byte内存空间里面的值
	var b *int
	b = &a //修改的是啥？8byte内存空间里面的值

	c := 30
	b = &c //b里面放的是c的地址

	//一次赋值多个内存块
	x, y, z := 10, "y", 6.43
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(x, y, z)
}
