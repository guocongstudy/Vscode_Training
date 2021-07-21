package main

import "fmt"

type Person1 struct {
	Name string
	Age  int
}

type Teacher struct {
	Person1
	TeacherId int
}

func main() {
	//声明了变量并且赋予零值
	var p1 Person1
	p1 = Person1{
		Name: "guocong",
		Age:  12,
	}
	var p2 = Person1{
		Name: "小李",
		Age:  13,
	}
	//看不用var的时候就需要用:=
	P3 := Person1{
		Name: "小刘",
		Age:  15,
	}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(P3)

	t1 := Teacher{
		Person1: Person1{
			Name: "gff",
			Age:  13,
		},
		TeacherId: 1,
	}
	fmt.Println(t1)
}
