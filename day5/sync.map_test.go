package day5

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"testing"
)

//-go 1.9引入内置方法,并发变成安全的map
//-sync.Map 将key和value 按照interface{}存储
//-查询出来后要用断言,x.(int) x.(string)
//-遍历使用Range()方法,需要传入一个匿名函数作为参数,匿名函数的参数为K,V interface{},每次调用匿名函数会将结果返回.
//可以搜索sync.map性能对比.

//-举例
func TestMap(t *testing.T) {
	m := sync.Map{}
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key_%d", i)
		m.Store(key, i)
	}
	//读取
	for i := 0; i < 15; i++ {
		key := fmt.Sprintf("key_%d", i)
		value, exist := m.Load(key)
		if exist {
			v := value.(int)
			log.Printf("[%s=%d]", key, v)
		} else {
			log.Printf("[key_not_found_in_sync.map:%s]", key)
		}
	}
	//删除
	m.Delete("key_9")
	//Range 遍历

	m.Range(func(k, v interface{}) bool {
		key := k.(string)
		value := v.(int)
		log.Printf("[找到了][%s=%d]", key, value)
		return true
	})
	log.Printf("return flase的遍历")
	m.Range(func(k, v interface{}) bool {
		key := k.(string)
		if strings.HasSuffix(key, "3") {
			log.Printf("不想要了3")
			return false
		} else {
			log.Printf("[找到了][%s=%d]", key, v.(int))
			return true
		}
	})
	//loadOrStore
	for i := 0; i < 15; i++ {
		key := fmt.Sprintf("key_%d", i)
		v, loaded := m.LoadOrStore(key, i)
		if loaded {
			//说明之前能有
			v := v.(int)
			log.Printf("[loadOrStore][之前有][%s=%d]", key, v)
		} else {
			//之前没有,新添加的
			log.Printf("[loadOrStore][之前没有,新添加的][%s=%d]", key, v.(int))
		}
	}
	value, Loaded := m.LoadAndDelete("key_11")
	log.Printf("[LoadAndDelete][key_11][v:%v][loaded:%v]", value, Loaded)
}

//1：37：50
//sync.Map使用场景
//-读多，给定key-v
