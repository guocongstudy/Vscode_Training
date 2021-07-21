package main

import "fmt"

func main() {
	a, _ := true, true //占位符，就是把值丢掉
	if !a {
		fmt.Println("flase")
	} else {
		fmt.Println("true")
	}
}
