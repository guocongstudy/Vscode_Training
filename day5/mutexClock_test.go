package day5

import (
	"log"
	"sync"
	"testing"
	"time"
)

/*
     go中的锁:
1.互斥锁
sync.mutex
--获取到锁的任务,阻塞其他任务
--意味着同一时间只有一个go才能持有锁

2.读写锁
sunc.ewmutex
*/

//举例:
var HcMutex sync.Mutex

//读写锁举例:
var rwMutex sync.RWMutex

//这个是互斥锁两个函数
func runMutex(id int) {
	//互斥锁就两种,锁和解锁
	log.Printf("[任务id:%d][尝试获取锁]", id)
	HcMutex.Lock()
	log.Printf("[任务id:%d][获取到了锁][开始干活:睡眠十秒]", id)
	time.Sleep(10 * time.Second)
	HcMutex.Unlock()
	log.Printf("[任务id:%d][干完活了 释放锁]", id)
}
func runHcLock() {
	//启动三个任务
	go runMutex(1)
	go runMutex(2)
	go runMutex(3)
}

//这个是读写锁两个函数
func runReadLock(id int) {
	//读锁
	log.Printf("[读任务:%d][进入读方法等待获取读锁]", id)
	rwMutex.RLock()
	log.Printf("[读任务:%d][获取到了读锁][干活,睡眠十秒]", id)
	time.Sleep(10 * time.Second)
	rwMutex.RUnlock()
	log.Printf("[读任务:%d][完成读任务,释放读锁]", id)
}

func runWriteLock(id int) {
	//读锁
	log.Printf("[写任务:%d][进入读方法等待获取写锁]", id)
	rwMutex.Lock()
	log.Printf("[写任务:%d][获取到了写锁][干活,睡眠十秒]", id)
	time.Sleep(10 * time.Second)
	rwMutex.Unlock()
	log.Printf("[写任务:%d][完成写任务,释放写锁]", id)

}

func allWriteWorks() {
	for i := 1; i < 6; i++ {
		//执行写任务
		go runWriteLock(i)
	}
}

func allReadWorks() {
	for i := 1; i < 6; i++ {
		//执行都任务
		go runReadLock(i)

	}
}

func writeFirst() {
	go runWriteLock(1)
	time.Sleep(1 * time.Second)
	go runReadLock(1)
	go runReadLock(2)
	go runReadLock(3)
	go runReadLock(4)
	go runReadLock(5)
}

func readFirst() {
	go runReadLock(1)
	go runReadLock(2)
	go runReadLock(3)
	go runReadLock(4)
	go runReadLock(5)
	time.Sleep(1 * time.Second)
	go runWriteLock(1)
}

//读写锁
//写锁阻塞所有锁,目的是修改时其他们不要读取,也不要修改
//读锁阻塞写锁,读锁可以同事的施加多个.目的:不要让修改数据影响我的读取结果.
//-同时多个读任务,可以施加多个读锁.阻塞写锁.
//-同时多个写任务,只可以施加一个写锁,阻塞其他所有锁,退化成互斥锁
//-读写混合:若有写锁,等待释放后能施加 读或者写
//-读写混合:若有写锁,等待释放后能施加,读或者写
func TestClock(t *testing.T) {
	//互斥锁
	//runHcLock()

	//同时多个写任务
	allWriteWorks()

	time.Sleep(6000 * time.Second)

}
