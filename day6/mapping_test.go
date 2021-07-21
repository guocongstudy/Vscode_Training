package main

import (
	"fmt"
	"log"
	"testing"
)

/*
反射的定义
-反射是指：类的应用，他们能够 自描述 自控制
-python反射 根据字符串执行函数，根据字符串导入包

#go中的反射
-go是静态语言，反射就是go提供的一种机制。在编译时不知道类型的情况下，可以做如下的事情
1.跟新变量
2.运行时查看值
3.调用方法
4.对他们的布局进行调整

##两个运用场景
1.你编写的这个函数，还不知道你的类型到底是什么，可能是还没约定好，也可能只传入的类型很多
2.希望通过用户的输入来决定这个函数（根据字符串调用方法），动态执行

举例使用 interface.type判断类型

##举例2 自定义struct的反射
1.对于成员变量
2.对于方法

*/

//func TestingMap(t *testing.T) {
//	var s interface{} = "abc"
//	switch s.(type) {
//	case string:
//		fmt.Println("这是一个string类型")
//
//	case int:
//		fmt.Println("这是一个int类型")
//
//	case float64:
//		fmt.Println("这是一个flaot64类型")
//	}
//
//}

type Person struct {
	Name string
	Age  int
}
type Student struct {
	Person     //匿名结构体嵌套
	StudentId  int
	SchoolName string
	IsBaoSong  bool
	Hobbies    []string
	hobbies    []string
	Labels     map[string]string
}

func (s *Student) GoHome() {
	log.Printf("[回家了][sid:%d]", s.StudentId)
}

func (s *Student) GotoSchool() {
	log.Printf("[上学了][sid:%d]", s.StudentId)
}

func (s *Student) baosong() {
	log.Printf("[保送了][sid:%d]", s.StudentId)
}

func TestMap1(t *testing.T) {
	s := Student{
		Person:     Person{Name: "xiaoyi", Age: 990},
		StudentId:  123,
		SchoolName: "汉中",
		IsBaoSong:  true,
		Hobbies:    []string{"唱", "跳", "rap"},
		hobbies:    []string{"唱", "跳", "rap"},
		Labels:     map[string]string{"k1": "v1", "k2": "v2"},
	}
	fmt.Println(s)
}
