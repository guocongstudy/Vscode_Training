package day5

import (
	"fmt"
	"log"
	"testing"
)

/*
##go的接口
-interface{}定义方法的集合
-多态体现在各个结构体对象，要实现接口中定义的方法。
-统一的函数调用入口，传入的接口
-各个结构体对象中，绑定的方法只能多不能少于接口定义的
-方法的签名要一致，参数类型，参数个数，方法名称，函数返回值要一致
-多态的灵魂就是有个承载的容器，先把所有实现了接口的对象加进来，每个实例都要照顾的地方，直接遍历，调用方法即可

*/
/*
1.多个数据源
2.query方法查数据
3.pushdata方法写入数据
*/

//方法集合
type DataSource interface {
	PushData(data string)
	QueryData(name string) string
}

type redis struct {
	Name string
	Addr string
}

func (r *redis) PushData(data string) {
	log.Printf("[Pushdata][ds.name:%s][data:%s]", r.Name, data)
}

func (r *redis) QueryData(name string) string {
	log.Printf("[Querydata][ds.name:%s][data:%s]", r.Name, name)
	return name + "redis"
}

type kafka struct {
	Name string
	Addr string
}

func (k *kafka) PushData(data string) {
	log.Printf("[Pushdata][ds.name:%s][data:%s]", k.Name, data)
}

func (k *kafka) QueryData(name string) string {
	log.Printf("[Querydata][ds.name:%s][data:%s]", k.Name, name)
	return name + "redis"
}

var Dm = make(map[string]DataSource)

func TestInterfacer(t *testing.T) {
	r := redis{
		Name: "redis",
		Addr: "1.1",
	}
	k := kafka{
		Name: "kafka",
		Addr: "2.2",
	}
	//注册数据源到承载的容器中
	Dm["redis"] = &r
	Dm["kafka"] = &k

	//推送数据
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key_%d", i)
		for _, ds := range Dm {
			ds.PushData(key)
		}
	}

	//查新数据
	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("key_%d", i)
		for _, ds := range Dm {
			res := ds.QueryData(key)
			log.Println(res)
		}
	}
}
