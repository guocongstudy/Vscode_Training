package main

import "fmt"

//ioto
const (
	A = iota //0
	B        //1
	c        //2
)
const (  //注意这里，不会变成3 而是从0重新开始
	X = iota //0
	Y        //1
	Z        //1
)

func main() {

	fmt.Println("hello,1")
	var array [2]int //[2]int{1,2}
	m := 0.01
	fmt.Println(m)
	fmt.Println(array)
	b := "strings" //根据变量类型，自动做了一个 var b string 的一个类型
	fmt.Println(b)

}
