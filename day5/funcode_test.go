package day5

import (
	"fmt"
	"testing"
)

/*
函数式编程
1.函数式编程的特点：
-高阶函数（函数参数和返回值都可以是函数）
-闭包
2.go函数编程的特点：
-主要是闭包


*/
//累加器
func commonAdd() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Printf("%d", sum)
}

//怎么使用函数式编程做这个累加器呢？
//Adder返回也是一个函数
func Adder() func(int) int {
	//自由变量
	num := 0
	return func(v int) int {
		num += v
		return num
	}
}

func callAddre() {
	addr := Adder()
	var res = 0
	for i := 0; i < 10; i++ {
		//整个的累加过程作为变量放在循环的外部
		//不断的对一个传入的数据进行加工
		res = addr(i)
		//在进行plus的加工
		fmt.Printf("+..%d=%d\n", i, res)
	}
}

func TestFuncode(t *testing.T) {
	//普通累计器
	commonAdd()
	//闭包累加器
	callAddre()
}

//3:11:33
