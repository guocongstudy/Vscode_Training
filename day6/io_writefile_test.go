package main

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestReader1(t *testing.T) {
	fileName := "go.txt"
	err := ioutil.WriteFile(fileName, []byte("升职加薪，迎娶白富美"), 0644)
	fmt.Println(err)

}
