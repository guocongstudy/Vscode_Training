package main

import (
	"log"
	"reflect"
	"testing"
)

/*
举例：反射调用方法
-过程说明
 1.首先reflect.Value(p1)获取得到反射类型对象
 2.reflect.ValueOf.MethodByName,需要传入准备的方法名称。MethodByName代表注册
 3.[]reflect.Value 这是最终调用方法的参数，无参数传空切片
/*反射的副作用
1.代码可读性变差
2.隐藏的错误要过编译检查
-go静态语言，编译能发现类型的错误
-但是对于反射代码是无能为力的，可能运行很久才会panic

反射的性能问题
type :=reflect.value(obj)
fieldValue:=type_,FieldByName("xx")

-每次取出fieldValue类型是reflect.value
-它是一个具体的值，不是一个可复用的对象
-每次反射都要malloc这个reflect.Value的结构体，还有GC

*/
type Person1 struct {
	Name   string
	Age    int
	Gender string
}

func (p Person) ReflectCallFuncWithArgs(name string, age int) {
	log.Printf("[调用的是带参数的方法][args.name:%s][args.age:%d][p.name:%s][p.age:%d]",
		name,
		age,
		p.Name,
	)
}

func (p Person) ReflectCallFuncWithNoArgs(name string, age int) {
	log.Printf("[调用的是不带参数的方法]")
}

func TestMapping1(t *testing.T) {
	p1 := Person1{
		Name:   "小乙",
		Age:    18,
		Gender: "男",
	}
	//1.首先通过reflectlite.ValueOf(p1)获取得到反射值类型
	getValue := reflect.ValueOf(p1)
	//2.参数方法调用
	methodValue := getValue.MethodByName("ReflectCallFuncWithArgs")
	//参数是reflect.Value的切片
	args := []reflect.Value{reflect.ValueOf("李逵"), reflect.ValueOf(30)}
	args = make([]reflect.Value, 0)
	methodValue.Call(args)
}

//1:48:54
