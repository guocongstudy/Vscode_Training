package main

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestReadDir(t *testing.T) {
	fs, err := ioutil.ReadDir("./")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, f := range fs {
		fmt.Printf("[name:%v][size:%v][isDir:%v][mode:%v][ModTime:%v]\n",
			f.Name(),
			f.Size(),
			f.IsDir(),
			f.Mode(),
			f.ModTime(),
		)

	}
}
