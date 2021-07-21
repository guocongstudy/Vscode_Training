package main

import (
	"fmt"
	"testing"
)

//go test -v scope_test.go  -v参数会显示每个用例的测试结果
func TestGlobal(t *testing.T) {
	a := 10
	b := 20
	fmt.Println(&a, &b)
	a = b
	fmt.Println(a, b)
	fmt.Printf("float32:%f", 3.14) //%f是小数但不含有指数。

	//字符串宽度控制
	fmt.Printf("|%s|", "aa")   //不设置宽度
	fmt.Printf("|%5s|", "aa")  //5个宽度，默认+，右对齐
	fmt.Printf("|%-5s|", "aa") //5个宽度，左对齐
	fmt.Printf("|%-5s|", "aa") //|000aa|

	//2.最大宽度，超出的部分会被截断
	fmt.Printf("|%.2s|", "xxx") //最大宽度为5



}
