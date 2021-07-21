package main

import (
	//"Vscode_demo/pkg" //这个顺序一定是正序，不能反着写，这个在放其他目录就没法运行，必须是在gopath/src中
	"Vscode_demo/pkg"
	"fmt"
)

const (
	MALE string = "MAN"
)

func add(x, y int) int {
	return x + y
}

//用于指明这个类型是一个接口
type Hello interface {
	Hello(username string) string
}

var (
	DefaultName = "郭聪"
)

const (
	//枚举类型
	a = iota //递增的这里的a从0开始
	//a = iota +1 不想从0开始，+1就是从1开始
	b //b=a+1
	c //c=b+1

)

func PointerTest(a string) string {
	a += "string b"
	return a
}

func PointerTest2(a *string) string {
	*a += "string b"
	return *a
}

type CoustB func(x, y int) int

func main() {
	a := 1 < 2 || 1 > 3
	fmt.Println(a)
	fmt.Printf(MALE)
	m1 := map[string]string{"username": "lao yu"}
	m := add(1, 2)
	m++
	b := "中午好"
	fmt.Println(b)
	//表示多行字符串，golang中是用反引号
	c := `guoocng 
dw dwd `
	fmt.Println(c)

	pkg.Demo()
	fmt.Println("hello words")
	fmt.Println(m)
	fmt.Println(m1)

	var d string
	d = "string a"
	fmt.Println(d, &d)
	d += "string b"
	fmt.Println(d, &d)

	f := "string a"
	fmt.Println(PointerTest(f))
	fmt.Println(f)

	g := "郭聪，好帅"
	//传的是指针对象，可以将地址里面的值改掉
	fmt.Println(PointerTest2(&g))
	fmt.Println(g)
	//将函数赋值给变量
	CoustB := func(x, y int) int { return 1 }
	fmt.Println(CoustB)
	//两种赋值方法
	//方法一：
	//var var1,var2,var3 float32
	//fmt.Printf("%d,%d,%d",var1,var2,var3)
	//方法二：
	var (
		var1 int
		var2 float32
		var3 string
	)
	var1 = 10
	var2 = 0.1
	var3 = "var3"
	fmt.Println(var1, var2, var3)

	var e string = "hello"
	fmt.Println(e)

}
