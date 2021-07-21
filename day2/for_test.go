package main

import (
	"fmt"
	"testing"
)

func TestForRange1(t *testing.T) {
	a := "dwdwddwf"
	for i, v := range a {
		fmt.Println(i, v)
		if v == 'f' {
			v = 'z'
			fmt.Println("change f")
		}
	}
	fmt.Println()
}
