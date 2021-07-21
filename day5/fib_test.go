package day5

import (
	"fmt"
	"testing"
)

/*函数式编程-斐波拉且数列*/
/*
#面向对象
-go 没有class关键字，oop，但是我们可以把go当作面向对象的方式来编程
-go中没有类，可以把struct作为类看待

#继承
-通过结构体的匿名嵌套，继承对应的字段和方法。

#多态
##go的接口
-interface{}定义方法的集合
-多态体现在各个结构体对象，要实现接口中定义的方法。
-统一的函数调用入口，传入的接口
-各个结构体对象中，绑定的方法只能多不能少于接口定义的
-方法的签名要一致，参数类型，参数个数，方法名称，函数返回值要一致

-实际应用prometheus的alert 和record



*/

type Person struct {
	Name  string
	Email string
	Age   int
}
type Stud struct {
	Person
	StudentId int
}

//附属于person类的方法
//指针相当于单实例绑定
func (p *Person) SayHello() {
	fmt.Printf("[person.SayHello][name:%s]", p.Name)
}
func Fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

//体现多态
//告警通知的函数，根据不同的对象通知
//有个共同的通知方法，每种对象自己实现

type notifer interface {
	//动作，定义的方法
	notify()
	query()
}

//普通用户
type user struct {
	name  string
	email string
}

//管理员
type admin struct {
	name string
	age  int
}

//在user和admin中实现
func (u *user) notify() {
	fmt.Printf("[普通用户][sendNotify to user %s\n]", u.name)
}
func (u *user) query() {

}
func (a *admin) notify() {
	fmt.Printf("[管理员][sendNotify to admin %s\n]", a.name)
}
func (a *admin) query() {

}

//多态的统一调用方法，入口
func sendNotify(n notifer) {
	n.notify()
}
func TestFib(t *testing.T) {
	Fib()
	p := Person{
		Name:  "xiaoyi",
		Email: "qq.com",
		Age:   18,
	}
	s := Stud{
		Person:    p,
		StudentId: 123,
	}
	s.SayHello()

	u1 := user{
		name:  "国雄",
		email: "dwdw",
	}
	a1 := admin{
		name: "力推",
		age:  18,
	}
	u1.notify()
	a1.notify()

	var n notifer
	n = &u1
	n.notify()

	ns := make([]notifer, 0)
	ns = append(ns, &u1)
	ns = append(ns, &a1)
	for _, n := range ns {
		//统一的入口
		n.notify()
	}
}
