package main

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestOperator3(t *testing.T) {
	s1 := "1"
	var i1 int8 = 1
	fmt.Println(unsafe.Sizeof(s1))
	fmt.Println(unsafe.Sizeof(i1))
	fmt.Printf("%b\n", 0b001)
}

//1:04:53
