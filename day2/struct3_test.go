package main

import (
	"fmt"
	"testing"
)

type Per struct {
	Name     string
	Age      int
	HouseIds [2]int
	Labels   map[string]string
}

func TestStruct3(t *testing.T) {
	p1 := Per{
		Name:     "小和",
		Age:      0,
		HouseIds: [2]int{32, 23},
		Labels:   map[string]string{"k1": "v1", "k2": "b2"},
	}
	p2 := p1
	fmt.Println(p2)

}

type test struct {
	name string
	age  int
	int  //匿名字段
}

/*为啥要有匿名字段，其实模仿面向对象中的继承*/
//深拷贝和浅拷贝
//02:51:04
//常量：例如，Π 全局用的时候不会被修改的。常量放在只读区域
