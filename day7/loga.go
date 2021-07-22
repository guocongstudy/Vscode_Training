package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	logger "github.com/sirupsen/logrus"
	"net/http"
)

type s struct {
	Name string
	Age  int
}

type dingHook struct {
	apiUrl     string //钉钉群机器人token_url
	levels     []logger.Level
	atMobiles  []string    //at谁
	appName    string      //模块前缀
	jsonBodies chan []byte //异步发送内容队列
	closeChan  chan bool   //主进程关闭消息通道
}

// -Levels 代表在哪几个级别下应用这个hook
func (dh *dingHook) level() []logger.Level {
	return dh.levels
}

// Fire -Fire代表 具体执行什么逻辑
func (dh *dingHook) Fire(e *logger.Entry) error {
	msg,_:=e.String()
	dh.DirectSend(msg)
	return nil
}

func (dh *dingHook) DirectSend(msg string) {
	dm := dingMsg{
		MsgType: "text",
	}
	dm.Text.Content = fmt.Sprintf("[日志告警log]\n[app=%s]\n"+"[日志详情：%s]", dh.appName, msg)
	dm.At.AtMobiles = dh.atMobiles
	bs, err := json.Marshal(dm)
	if err != nil {
		logger.Errorf("[消息json.marshal失败][error:%v][msg:%v]", err, msg)
		return
	}
	res, err := http.Post(dh.apiUrl, "application/json", bytes.NewBuffer(bs))
	if err != nil {
		logger.Errorf("[消息发送失败][error:%v][msg:%v]", err, msg)
		return
	}
	if res != nil && res.StatusCode != 200 {
		logger.Errorf("[钉钉返回错误][error:%v][msg:%v]", res.StatusCode, msg)
		return
	}
}

type dingMsg struct {
	MsgType string `json:"msgType"`

	Text struct {
		Content string `json:"content"`
	} `json:"text"`

	At struct {
		AtMobiles []string `json:"atMobiles"`
	} `json:"at"`
}

func main() {

}
