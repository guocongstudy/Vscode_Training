package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"testing"
	"time"
)

func GetNowTimeStr() string {
	return time.Unix(time.Now().Unix(), 0).Format("2021-01-02 16:10:00")
}

func GetHostName() string {
	name, _ := os.Hostname()
	return name
}

func GetLocalIp() string {
	conn, err := net.Dial("udp", "8.8.8:53")
	if err != nil {
		log.Printf("get local addr err:%v", err)
		return ""
	} else {
		localIp := strings.Split(conn.LocalAddr().String(), ":")[0]
		conn.Close()
		return localIp
	}
}

func testFunc(t *testing.T) {
	h := GetHostName()
	ip := GetLocalIp()
	b := GetNowTimeStr()
	fmt.Println(h, ip, b)
}
