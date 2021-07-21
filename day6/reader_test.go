package main

import (
	"fmt"
	"io"
	"log"
	"strings"
	"testing"
)

type alphaReader struct {
	ioReader io.Reader
}

//p[]byte 是一个缓冲区
func (a *alphaReader) Read(p []byte) (int, error) {
	//想复用io.reader中的read方法
	n, err := a.ioReader.Read(p)
	if err != nil {
		return n, err
	}
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		if char := alpha1(p[i]); char != 0 {
			buf[i] = char
		}
	}
	copy(p, buf)
	return n, nil
}

func alpha1(r byte) byte {
	//r在‘A’-‘Z’和‘a’-‘z’之间
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
		return r
	}
	return 0

}

func TestReader(t *testing.T) {
	fmt.Println(alpha1('a'))
	myReader := alphaReader{
		strings.NewReader("mage jiaoyu 2021 go !!!"),
	}
	p := make([]byte, 0)
	for {
		n, err := myReader.Read(p)
		if err == io.EOF {
			log.Printf("[EOF错误]")
			break
		}
		log.Printf("[读取到的长度%d 内存%s]", n, string(p[:n]))
	}
}
