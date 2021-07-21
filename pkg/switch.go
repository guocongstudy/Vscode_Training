package pkg

import (
	"fmt"
	"testing"
)

func ExampleSwitch() {
	grade := "A"
	switch {
	//case条件后面不能用中括号，要用冒号
	case grade == "A":
		fmt.Println("优秀呀")
	case grade == "B":
		fmt.Println("差不多呀")
	case grade == "C":
		fmt.Println("不行呀 你")
		//默认输出
	default:
		fmt.Println("好好学习，你可以的")
	}
}

func TestSwitch(t testing.T) {

}
