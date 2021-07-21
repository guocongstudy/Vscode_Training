package main

import (
	"fmt"
	"testing"
	"unsafe"
)

//切片
type slice struct {
	array unsafe.Pointer //数组指针
	len   int            //长度
	cap   int            //容量
}

func TestSlice(t *testing.T) {
	//初始化一个长度为3容量为4的切片
	a := make([]int, 3, 5)
	fmt.Println(a, len(a), cap(a))
	fmt.Println(a[2])
	fmt.Println(a, len(a), cap(a))

	//初始化 5的位置放的值为100
	s1 := []int{5: 100}
	fmt.Println(s1, len(s1), cap(s1))
	//切片和数组的区别
	/*    1.数组
	          a :=[3]int{1,2,3} 返回的是值
		      2.切片
		      b :=[]int{1,2,3}  返回的是地址
	*/
	var s2 []int
	fmt.Printf("%p\n", s2)
	s2 = make([]int, 0)
	fmt.Printf("%p\n", s2)
	//数组是不可以扩容的
	s2 = append(s2, 1, 2, 3, 4)
	fmt.Println(s2)
	fmt.Println(s2, len(s2), cap(s2))

}

func TestSlice1(t *testing.T) {
	slice := []int{3: 100}
	fmt.Println(slice, len(slice), cap(slice))
	s1 := make([]int, 0, 4)
	s1 = append(s1, 3, 4, 22, 42)
	//将s1赋值给s2
	s2 := s1[:]
	fmt.Println(s1, s2)

	s2[0] = 100
	fmt.Println(s1[0], s2[0])

	for i, v := range s1 {
		fmt.Println(i, v)
	}

	s3 := make([]int, 0, 4)
	s3 = append(s3, 10, 20, 30, 40)
	//左闭右开
	s4 := s3[1:3]
	fmt.Println(s4)

}
