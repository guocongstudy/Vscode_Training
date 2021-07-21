package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

/*


 */
func numLen(n int64) int {
	return len(strconv.Itoa(int(n)))
}

func main() {
	now := time.Now()
	log.Printf("[当前时间对象为：%v]", now)
	log.Printf("[当前时间戳秒级：%v][位数：%v]", now.Unix(), numLen(now.Unix()))
	//返回日期
	fmt.Println(now)
	year, month, day := now.Date()
	log.Printf("[通过now.Data获取][年：%d 月：%d 日：%d]", year, month, day)
	log.Printf("[直接获取年：%d]", now.Year())
	log.Printf("[直接获取月：%d]", now.Month())
	log.Printf("[直接获取日：%d]", now.Day())
	//时。分。秒也可以这么输出。
	log.Printf("[全部：%v]", now.Format("2006-01-02 15:04:05"))

	tsr := "2021-07-17 16:52:58"
	layout := "2006-01-02 15:04:05"

	//Parse：将字符串格式的时间转为 time.Time
	//time.Local 指的是时区
	t1, _ := time.ParseInLocation(layout, tsr, time.Local)
	t2, _ := time.ParseInLocation(layout, tsr, time.UTC)
	log.Printf("[%s的 UTC时区的时间戳为：%d]", tsr, t1.Unix())
	log.Printf("[%sde  UTC时区的时间戳为：%d]", tsr, t2.Unix())
	log.Printf("[UTC -cst =%d]", t2.Unix()-t1.Unix())

}
