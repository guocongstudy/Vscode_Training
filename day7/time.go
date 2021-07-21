package main

import (
	"log"
	"time"
)

var layout = "2021-07-02 18:21:00"

func tTostr(t time.Time) string {
	return time.Unix(t.Unix(), 0).Format(layout)
}

//日期比较
func main() {
	now := time.Now()
	t1, _ := time.ParseDuration("1h")
	m1 := now.Add(t1)
	log.Printf("[a.after(b) a在b之后：%v]", m1.After(now))
	log.Printf("[a.Before(b) a在b之前：%v]", m1.Before(now))
	log.Printf("[a.after(b) a在b之后：%v]", m1.After(now))
}
