package main

import (
	"log"
	"sync"
	"testing"
	"time"
)

/*
-interfce
-发布系统，开启新的k8s，node（jtype字段）不同类型的任务
-增量更新，开启新的，停掉旧的
-a,b,c    b,c,d -->开启d，停掉a
*/

//定义他的接口
type deployJob interface {
	//有新的元素就开启
	start()
	//本地有的，但是远程控制端没有的
	stop()
	//job唯一的判定是一个哈希
	hash() string
}

type jobManager struct {
	targetMtx     sync.RWMutex
	activeTargets map[string]deployJob
}

func (jm *jobManager) sync(jobs []deployJob) {
	//增量更新体现的是：新的更新，旧的删除
	//在job里面，不在activeTargets 说明新增
	//在job里面，也在acticeTargets 说明不动
	//不在job里面，在activeTargetds 说明删除
	thisAll := make(map[string]deployJob)
	thisNew := make(map[string]deployJob)
	jm.targetMtx.Lock()
	//第一次遍历结束，拿到all和new
	for _, t := range jobs {
		hash := t.hash()
		thisAll[t.hash()] = t
		if _, loaded := jm.activeTargets[hash]; !loaded {
			thisNew[hash] = t
			jm.activeTargets[hash] = t
		}
	}
	//判断旧的
	for hash, t := range jm.activeTargets {
		if _, loaded := thisAll[hash]; !loaded {
			//先把旧的任务停掉
			t.stop()
			//更新一下
			delete(jm.activeTargets, hash)
		}
	}

}

type k8sD struct {
	Id    int
	Name  string
	k8sNS string
}

type hostD struct {
	Id     int
	Name   string
	HostIp string
}

func (kd *k8sD) start() {
	log.Printf("[k8s.delpoy.start][%v]", kd)
}

func (kd *k8sD) stop() {
	log.Printf("[k8s.delpoy.stop][%v]", kd)
}

func (kd *k8sD) hash() string {
	return kd.Name
}

func (h *hostD) start() {
	log.Printf("[host.delpoy.start][%v]", h)
}

func (h *hostD) stop() {
	log.Printf("[host.delpoy.stop][%v]", h)
}

func (h *hostD) hash() string {
	return h.Name
}

func TestHardword2(t *testing.T) {
	time.Sleep(10 * time.Second)

}

//pkg
//58:59
