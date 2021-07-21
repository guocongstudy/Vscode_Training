package courseware_test

import (
	"fmt"
	"testing"
)

/*
写的时候也发现两个问题：
1.写测试函数的时候，函数名用Test 而不能用Testing
2.go test -v addThree_test.go   这个里面的-v是小写，不能是大写
*/
func TestAddThree(t *testing.T) {
	op := []int{1, 2, 3, 4, 5}
	target := AddThree(op, func(x int) int {
		return x + 3
	})
	fmt.Println(target)
}

//4,5,6,7,8
type Functor func(x int) int

func AddThree(operand []int, fn Functor) (target []int) {
	for _, v := range operand {
		target = append(target, fn(v))
	}
	return
}
