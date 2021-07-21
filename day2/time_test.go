package main

import (
	"fmt"
	"testing"
	"time"
	"unsafe"
)

func TestTime(t *testing.T) {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Format("2008-01-02 15:04"))
	fmt.Println(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	timeStr := "2021-06-19 16:44:14"
	t1, err := time.Parse(time.RFC3339, timeStr)
	fmt.Println(t1, err)

}

func TestTime1(t *testing.T) {
	now := time.Now()
	fmt.Println(now.Unix())

	ts := 2424343
	t1 := time.Unix(int64(ts), 0)
	t1.Format("2002-02-21")
}

//不安全指针
func TestUnsafePointer(t *testing.T) {
	a := "string a"
	p1 := &a //*string string的指针
	fmt.Printf("%p\n", p1)

}

//go中不能直接对地址进行操作，为了就是防止不安全的行为，如果项目中真的还有这种需求，就要用到unsafe标准库、
func TestPonter(t *testing.T) {
	var x int64 = 20
	a := &x
	fmt.Println(a)

}

type Man struct {
	Name string
	Age  int64
}

func TestUnsafePonter1(t *testing.T) {
	m := Man{Name: "Join", Age: 20}
	fmt.Println(unsafe.Sizeof(m.Name), unsafe.Sizeof(m.Age), unsafe.Sizeof(m))
	fmt.Println(unsafe.Offsetof(m.Name))
	fmt.Println(unsafe.Offsetof(m.Age))
}
