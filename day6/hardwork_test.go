package main

import (
	"fmt"
	"sync"
	"testing"
)

//Go语言包与工程化
//简单的作业
type SafeSet struct {
	sync.RWMutex
	m map[string]struct{}
}

func NewSet() *SafeSet {
	return &SafeSet{
		m: make(map[string]struct{}),
	}
}

func (ss *SafeSet) Add(key string) {
	ss.Lock()
	defer ss.Unlock()
	ss.m[key] = struct{}{}
}

func (ss *SafeSet) Del(key string) {
	ss.Lock()
	defer ss.Unlock()
	delete(ss.m, key)
	ss.m[key] = struct{}{}
}

func (ss *SafeSet) JudeElement(key string) bool {
	ss.RLock()
	defer ss.RUnlock()
	_, ok := ss.m[key]
	return ok
}

func (ss *SafeSet) PrintElement() []string {
	ss.RLock()
	defer ss.RUnlock()
	res := make([]string, 0)
	for k := range ss.m {
		res = append(res, k)
	}
	return res
}

func (ss *SafeSet) Merge(set *SafeSet) {
	ss.Lock()
	defer ss.Unlock()
	keys := set.PrintElement()
	for _, k := range keys {
		//遍历的时候还没有对写锁解锁
		//这时不能调用ss.Add,因为Add也要加写锁，锁中锁的问题
		ss.m[k] = struct{}{}
	}
}

func SetAdd(set *SafeSet, n int) {
	for i := 0; i < n; i++ {
		key := fmt.Sprintf("key_%d", i)
		set.Add(key)
	}
}

func TestHardWork(t *testing.T) {
	s1 := NewSet()
	SetAdd(s1, 10)
	res := s1.PrintElement()
	fmt.Println(res)

}
