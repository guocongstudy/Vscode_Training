package main

import (
	"fmt"
	"os"
	"testing"
)

/*
##os.Create
-创建得到的file对象可以用，write和writestring写内容
-


*/
func TestCreate(t *testing.T) {
	file, err := os.Create("a.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	for i := 0; i < 10; i++ {
		file.WriteString("writeString写进来的\n")
		file.Write([]byte("wrtie写进来的\n"))

	}
}
