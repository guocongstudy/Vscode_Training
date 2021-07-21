package main

import (
	"fmt"
	"log"
	"os"
)

//日志管理
func logPrinta(baseStr string) {
	for i := 0; i < 10; i++ {
		msg := fmt.Sprintf("%s_%d", baseStr, i)
		log.Println(msg)
	}
	//
}

func main() {
	//创建文件对象
	file, err := os.OpenFile("my.log", os.O_APPEND|os.O_APPEND|os.O_WRONLY, 06666)
	if err != nil {
		log.Fatal(err)
	}
	//设置log输出到文件
	log.SetOutput(file)
}

//3:25:57
