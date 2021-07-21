package main

import (
	"fmt"
	"testing"
)

func TestTT(t *testing.T) {
	fmt.Println(TT())
}

func TT() func() int {
	x, y := 10, 20
	fmt.Println(x, y)
	add := func() int {
		return x + y
	}
	return add

}
