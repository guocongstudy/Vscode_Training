package main

import (
	"fmt"
	"testing"
)
//
///*swithc默认情况下case最后自带break语句，匹配成功后就不会执行其他case，如果我们需要执行后面的case，可以使用fallthrough*/
//func main() {
//	caseB := true
//	switch caseB {
//	//用break跳出以后，不会再执行后面的语句，但是可以用fallthrough去替换，可以将下面的语句继续执行
//	//下面执行的还是条件等同的语句
//	case true:
//		fmt.Println("case 1")
//		fallthrough
//	case false:
//		fmt.Println("case 2")
//		fallthrough
//		//看这里应该是软件的问题，能运行但是还是飘红。
//	case true:
//		fmt.Println("case true")
//		//fallthrough 穿透
//		fallthrough
//		//当上面所有问题都不满足的时候，就会执行fallthrough语句。
//	default:
//		fmt.Println("default")
//	}
//}

/*
1.又get到一个点，自己写了Test函数，但是如果文件夹不加“_test”就不会出现test的运行按钮，函数名也不会变亮
2.要开始一个单元测试，需要准备一个 go 源码文件，在命名文件时需要让文件必须以_test结尾。默认的情况下，go test 命令不需要任何的参数，它会自动把你源码包下面所有 test 文件测试完毕，当然你也可以带上参数。

这里介绍几个常用的参数：

    -bench regexp 执行相应的 benchmarks，例如 -bench=.；
    -cover 开启测试覆盖率；
    -run regexp 只运行 regexp 匹配的函数，例如 -run=Array 那么就执行包含有 Array 开头的函数；
    -v 显示测试的详细命令

*/
//运行测试文件，命令为：go test -v fallthrough_test.go
func TestSwitchFallthrough(t *testing.T){
	a :=10
	switch a {
	case 10:
		fmt.Println("10")
		fallthrough
	case 20:
		fmt.Println("20")
	default:
		fmt.Println("default")
	}
}