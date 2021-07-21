package day5

import (
	"fmt"
	"testing"
)

func TestPointer1(t *testing.T) {
	//指针的操作interesting
	a := new(int)
	fmt.Println(&a, a, *a)

	//& 取变量的内存地址
	//* 读取地址的值

}

type Student struct {
	Name     string   //名字
	Number   uint16   //学号
	Subjects []string //数学,语文,英语
	Score    []int    //88,99,77
}

type Class struct {
	Name     string     //班级名称
	Number   uint8      //班级编号
	Subjects []string   //数学,语文,英语
	students []*Student //班级学员,student的指针的切片
}

//给结构体写个方法,Class的方法
func (c *Class) AvgScore() []int {
	//遍历切片
	//获取每一个学院的第一个成就,数据成绩
	return nil
}
