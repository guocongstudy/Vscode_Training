package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

/*
#ioutil库，工具包
-在io目录中，它是一个工具包，实现一些实用的工具
-

*/

func TestIo1(t *testing.T) {
	fileName := "go.txt"
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Printf("[内容：%s]", bytes)
}
