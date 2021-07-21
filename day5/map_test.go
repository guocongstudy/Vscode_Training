package day5

import (
	"sync"
	"testing"
)

func TestMap1(t *testing.T) {
	normalM := map[string]string{
		"guocong":   "122",
		"zhangxu":   "232",
		"zhangheng": "232",
	}
	mm := sync.Map{}
	for k, v := range normalM {
		mm.Store(k, v)
	}
}

/*结构体复杂的case多不用sync.map*/
//func TestMap01(t *testing.T) {
//	m := cmap.New()
//	//写map的 go
//	go func(){
//		for i:=0;i<1000;i++{
//			key :=fmt.Sprintf("key_%d",i)
//			m.Set(key,i)
//		}
//	}
//}
