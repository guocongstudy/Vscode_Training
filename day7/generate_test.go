package main

import (
	"log"
	"time"
)

func main() {
	now := time.Now()
	log.Printf("[当前时间对象为：%v]", now)
	log.Printf("[当前时间戳秒级为：%v]", now)
}
