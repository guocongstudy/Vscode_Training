package main

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {

	var slice = make([]int, 0, 5)
	fmt.Println("没有添加元素之前，slice的样子:")
	fmt.Println(slice)
	fmt.Println("开始添加元素，slice的长度和容量变化:")
	b := append(slice, 1)
	fmt.Println(b)
	fmt.Printf("长度为：%d，容量为：%d\n", len(b), cap(b))
}
