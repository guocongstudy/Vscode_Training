package main

import (
	"fmt"
	"testing"
)

/*
#单元测试
1.什么是单元测试
-单元测试是应用最小的可测试部件，如函数对象的方法
2.什么是单元和单元测试？
单元测试是软件开发对最小单元进行正确性检验的测试工作
3.为什么进行单元测试？
-保证 变更/重试的正确性，特别是一些频繁变动和多人合作开发的项目中
-简化调试过程，可以轻松的让我们知道那一部分代码出现问题
-单元测试也是最好用的文档，在单测中会直接给出具体接口的使用方法

#单元测试用例编写的原则
1.单一原则：一个测试用例只赋值一个场景
2.原子性：结果只有两种情况：pass/Fail
3.优先要写核心组件和逻辑的测试用例
4.高频使用库，util 终点覆盖这些

#单测用例约定
1.文件名必须要xx_test.go命名
2.测试方法必须是Testxxx开头
3.方法中的参数必须是t*testing.T
4.测试文件和被测文件必须是在同一包下，不能跨包测

#golang的测试框架
##testing
go test -v
只想执行测试函数中的一个函数？

go test -run=TestAdd -v .
正则过滤函数名？
go test -run=TestM.* -v .
(只执行以M开头的测试函数)


##-cover测试覆盖率
-用于统计目标包百分之多少的代码参与了测试？
go test -v -cover
-用于看某一个函数的覆盖率？
go test -v -cover -run=TestAdd

#table-drivern tests
1.所有的用例的数据组织在切片cases中，看起来就像一张表，借助循环创建子测试，这样写的好处有：
 -新增用例非常简单，只需要给cases新增一条测试数据即可
 -测试代码可读性好，直接能够看到每个子测试的参数和期待的返回值
 -用例失败时，报错信息的格式比较统一，测试报告易于阅读
 -如果数据量较大，或者是一些二进制数据，推荐使用相对路径从文件中读取。

#基准测试


##子测试 t.run

##table-driven tests

*/

type Str struct {
	Name    string
	Chinese int
	Math    int
}

func NewStudent(name string) (*Str, error) {
	if name == "" {
		return nil, fmt.Errorf("name为空")
	}
	return &Str{
		Name: "郭聪",
	}, nil
}

func (s *Str) GetAvgScore() (int, error) {
	score := s.Chinese + s.Math
	if score == 0 {
		fmt.Println("是零分大哥！牛逼")
	}
	fmt.Println("你的平均分是:")
	return score / 2, nil

}

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-2) + fib(n-1)
}

//基准测试
//go test -bench=. -benchmem -run=none
func BenchmarkFib(b *testing.B) {

}
