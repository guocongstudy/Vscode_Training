package main

import "fmt"

func main() {
	fmt.Printf("%T,%T\n", a, b)
	fmt.Println()
}

//函数不能进行比较
func a(a int) {
	fmt.Println(a)
}

func b(b int) {
	fmt.Println(b)
}
