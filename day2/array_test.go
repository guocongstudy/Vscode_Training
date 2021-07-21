package main

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	//定义一个数组，长度为四，类型为int
	var a [4]int                 //8B*4 32byte
	var b [4]float64             //B*4 4Byte
	var c = [4]int{1, 23, 34, 4} //不能用:=
	fmt.Println(c)
	fmt.Println(a)
	fmt.Println(b)

	//声明一个指针数组
	var d = [4]*int{0: new(int), 2: new(int)}
	v := 10
	d[0] = &v
	d[1] = &v
	fmt.Println(d)
	//将指针的地址转成值
	fmt.Println(*d[0], *d[1])
}

func TestArray2(t *testing.T) {
	//数组的遍历，for range
	m := new(int)
	fmt.Println(m)
	var a = [5]int{1, 23, 4, 5, 56}
	for i, v := range a {
		fmt.Println(i, v)
	}
	//二维数组
	b := [4][4]int{{1, 1, 1, 1}, {1, 2, 2, 2}, {1, 2, 2, 3}, {1, 4, 3, 2}} //(0,0),(0,1),(0,2),(0,3),(0,4)
	fmt.Println(b)

}

func change(payload [4]int) {
	payload[0] = 10
}

type Person struct {
	Name          string
	Id            int
	Address       string
	FavoriteColor []string
	Addr          Home
}

//前面的（person）指明是属于person的方法， p为别名
func (p Person) add(x, y int) int {
	return p.Id * 2
}

func TestMain2(t *testing.T) {
	//var person Person
	//加了冒号就不用去var
	person := &Person{}
	fmt.Println(person)
}

func TestPerson(t *testing.T) {
	//记住var 自己写很容易将var写掉
	var person = Person{
		Name:          "guocong",
		Id:            12,
		Address:       "汉中",
		FavoriteColor: []string{"blue", "black"},
	}
	fmt.Println(person)
	fmt.Println(person.Name)
	fmt.Printf("%d\n", person.Id)

	m := new(Person)
	m.Addr.City = "汉中"
	m.Id = 6
	fmt.Printf("%T\n", m)
	fmt.Println(m.add(4, 65))
	fmt.Println(m)
}

type Home struct {
	City string
}

type Author struct {
	Name string
	Age  int
}

func (a *Author) GetName() string {
	return a.Name
}

type Title struct {
	Main string
	Sub  string
}

type Address struct {
	Address string
}

type Book struct {
	Author Author
	Title  Title
	//匿名嵌套
	Address
}

func TestBook(t *testing.T) {
	b := &Book{
		Author: Author{
			Name: "中华上下五千年",
			Age:  12,
		},
		Title: Title{
			Main: "t1",
			Sub:  "t2",
		},
		//匿名嵌套，还是用原来的变量
		Address: Address{
			Address: "汉中",
		},
	}
	fmt.Println(b)
}
