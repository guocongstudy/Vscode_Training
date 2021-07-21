package day5

import (
	"fmt"
	"testing"
)

//变量是一个地址，也就是一个引用，如果我们想要保持这个内存地址怎么办

func TestPointer(t *testing.T) {
	//指针的演化
	a := 10
	fmt.Println(&a)
	c := &a
	fmt.Println(c)
	fmt.Printf("%T\n", c)
	fmt.Println(&c)

	var e *int
	var f *float32
	fmt.Println(e, f)
	e = new(int) //将返回的地址存入e中
	fmt.Println(a)

	//new不会初始化,new开闭一个类型对应的内存空间,返回一个内存空间的地址

	//make,返回的也是一个内存地址,

}
