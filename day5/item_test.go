package day5

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

type item struct {
	value int
	ts    int64 //时间戳，要么是被创建出来的时间，要么是被更新出来的时间
}

type Cache struct {
	sync.RWMutex
	mp map[string]*item
}

func (c *Cache) Get(key string) *item {
	c.RLock()
	defer c.RUnlock()
	return c.mp[key]

}

func (c *Cache) Set(key string, value *item) {
	c.Lock()
	defer c.Unlock()
	c.mp[key] = value
}

func (c *Cache) Gc(timeDelta int64) {
	for {
		c.Lock()
		now := time.Now().Unix()
		//遍历缓存中的项目，对比时间戳，超过timeDelta的删除
		for k, v := range c.mp {
			if now-v.ts > timeDelta {
				log.Printf("[这个项目过期了][key %s]", k)
				delete(c.mp, k)
			}
		}
		c.Unlock()
		time.Sleep(5 * time.Second)
	}
}

func TestMain1(t *testing.T) {
	c := Cache{
		mp: make(map[string]*item),
	}
	//让删除过期任务异步执行，异步执行
	go c.Gc(30)
	//写入数据,从mysql读取。
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key_%d", i)
		ts := time.Now().Unix()
		im := &item{
			value: i,
			ts:    ts,
		}
		//设置缓存
		log.Printf("[设置缓存][项目][key:%s][v:%v]", key, im)
		c.Set(key, im)
	}
	time.Sleep(31 * time.Second)

}
