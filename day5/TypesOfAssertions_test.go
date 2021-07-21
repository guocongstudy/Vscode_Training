package day5

import (
	"fmt"
	"testing"
)

/*
##空接口

##类型断言和类型判断
- 一个interface需要类型转换，语法 i.(T)
- v,ok
*/

func TestInterface(t *testing.T) {

	var s interface{} = "abc"
	s1, ok := s.(string)
	fmt.Println(s1, ok)

	s2, ok := s.(int)
	fmt.Println(s2, ok)

	switch s.(type) {
	case string:
		fmt.Println("是string")
	case int:
		fmt.Println("是int")
	case float64:
		fmt.Println("是float")
	}

}
