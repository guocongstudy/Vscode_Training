package day5

import "sync"

//读写锁
type concurrentMap struct {
	sync.RWMutex
	mp map[int]int
}

func (c *concurrentMap) Set(key int, value int) {
	//先获取锁
	c.Lock()
	//set值
	c.mp[key] = value
	//解锁
	c.Unlock()
}

func (c *concurrentMap) Get(key int) int {
	//先获取读锁
	c.RLock()
	//获取值
	res := c.mp[key]
	//解锁
	c.RUnlock()
	return res
}
