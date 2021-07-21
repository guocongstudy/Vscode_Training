package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

/*
上述都提供了文件读写的能力
-bufio多了一层缓存的能力，优势体现在读取大文件的时候
-ioutil.ReadFile是一次性将内容加载到内存，大文件容易爆掉

##

*/

func TestRead(t *testing.T) {
	fileName := "a.txt"
	//方式一
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	log.Printf("[方法一 ioutil.ReadFile][res:%v]", string(bytes))

	//
	file, err := os.Open(fileName)
	if err != nil {
		return
	}
	bytes, err = ioutil.ReadAll(file)
	if err != nil {
		return
	}
	log.Printf("[方法二 os.Open+ioutil.ReadAll][res:%v]", string(bytes))
	file.Close()

	file, _ = os.Open(fileName)
	buf := make([]byte, 50)
	_, err = file.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Printf("[方法三 os.Open+file.Read][res:%v]", string(bytes))
	file.Close()

	file, _ = os.Open(fileName)
	//bufio.NewReader
	rd := bufio.NewReader(file)
	buf1 := make([]byte, 50)
	_, err = rd.Read(buf1)
	if err != nil {
		return
	}
	log.Printf("[方法四 os.Open+bufio.Read][res:%v]", string(bytes))
	file.Close()
}
