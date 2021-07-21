package main

import (
	"io"
	"log"
	"testing"
)

/*
Go输入和输出操作是使用原语言实现的
#IO操作也叫输入输出操作
-IO操作主要是用来读取或写入数据，在很多语言中也叫作流操作

#Go输入和输出操作是使用原语实现的
-这些原语将数据模拟成可以读可以写的字节流

#IO库属于底层接口定义库
-作用是定义一些基本的接口和基本常量
-其他人io.EOF

#IO库比较常用的接口有三个
-分别是Reader，Writer和Close。
-IO库中实现的上述接口可以以流的方式处理数据
-并不需要考虑数据是什么？数据来自哪里？数据发送到哪里？

##Reader
-io.Reader表示一盒读取器
-它从某个地方读取数据到缓存区
-在缓存区里面存储可以流式的使用。
-接口签名如下：
  type Reader interface{
    //[]byte相当于是一个读缓冲区
    Read(p []byte)(n int,err error)
-Read()首先要有一个读缓冲区的参数
-Read()返回两个值，第一个是读取到的字节数，第二个读取时发生的错误。
-注意：返回的读取字节个数n可能小于缓冲区的大小
-io.EOF表示的输入的流已经读到头了

##自己实现一个reader
-要求过滤输入字符串中的非字母字符

##组合多个Reader，目的是重用和屏蔽下层实现的复杂度
-标准库里面已经有很多Reader
-使用一个Reader A作为一个Reader B的一部分
-目的是复用逻辑，流式处理
-复用的io.Reader
}
#
*/

type zimuguotv struct {
	src string //输入的空字符串
	cur int
}

func alpha(r byte) byte {
	//字母在A-Z，a-z之间
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
		return r
	}
	return 0
}

func (z *zimuguotv) Read(p []byte) (int, error) {
	//当前位置>=字符串的长度，说明读取到结尾了,返回EOF
	if z.cur >= len(z.src) {
		return 0, io.EOF
	}
	//定义一个剩余还没读到的长度
	//总长度-当前读到的位置
	x := len(z.src) - z.cur
	//bound叫做本次读取的长度
	//n代表本次遍历bound的大小
	n, bound := 0, 0
	if x >= len(p) {
		//剩余长度超过缓冲区大小，说明本次可以完全填满缓冲区
		bound = len(p)
	} else {
		//剩余长度小于缓冲区大小，剩余长度进行输出，缓冲区填不满。
		bound = x
	}
	buf := make([]byte, bound)
	for n < bound {
		if char := alpha(z.src[z.cur]); char != 0 {
			buf[n] = char
		}

		//n++代表索引++
		n++
		z.cur++
	}
	//将本次得到的数据拷贝给p
	copy(p, buf)
	return n, nil
}

func TestBuf(t *testing.T) {
	zmreader := zimuguotv{
		src: "mage jiaoyu 2021 g@ !!!",
	}
	p := make([]byte, 4)
	for {
		n, err := zmreader.Read(p)
		if err == io.EOF {
			log.Printf("[EOF错误]")
			break
		}
		log.Printf("[读取到的长度%d,内存%s]", n, string(p[:n]))
	}
}

//func TestIo(t *testing.T) {
//	//实现一个redaer每次读取4个字节
//	//从字符串创建一个reader对象
//	reader := strings.NewReader("马哥教育 2021 golang")
//	//new一个4字节的读取缓冲区
//
//	p := make([]byte, 4)
//	for {
//		//reader对象怎么读数据呢？
//		n, err := reader.Read(p)
//		if err != nil {
//			if err == io.EOF {
//				fmt.Printf("[数据已经读完，Eof:%d]", n)
//				break
//			}
//			log.Printf("[未知错误：%v]", err)
//			return
//		}
//		log.Printf("[打印读取的字节数：%d 内容：%s]", n, string(p[:n]))
//	}
//}
