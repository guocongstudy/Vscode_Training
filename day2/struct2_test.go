package main

import (
	"Vscode_demo/tt"
	"fmt"
)

type Person2 struct {
	name string
	Age  int
}

func (p Person2) GetName() string {
	return p.name
}

func main() {
	fmt.Println(tt.Abc)
}
