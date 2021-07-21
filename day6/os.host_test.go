package main

import (
	"log"
	"os"
	"testing"
)

func TestHost(t *testing.T) {
	log.Printf("[获取命令行参数][res:%v]", os.Args)
	//获取主机名
	hn, _ := os.Hostname()
	log.Printf("[获取主机名][res:%v]", hn)
	log.Printf("[获取主机名][res:%v]", os.Getegid())
	log.Printf("[获取一条环境变量][res:%v]", os.Getenv("GOROOT"))
	//获取所有环境变量
	env := os.Environ()
	for _, v := range env {
		log.Printf("[获取所有环境变量][res:%v]", v)
	}
	dir, _ := os.Getwd()
	log.Printf("[获取当前目录][res:%v]", dir)

	_ = os.Mkdir("config", 0755)
	log.Printf("[创建单一目录config目录]")

	//mkdir -p
	os.MkdirAll("config1/yaml/local", 0755)
	log.Printf("[递归创建目录config目录]")

	os.Remove("config")
}
