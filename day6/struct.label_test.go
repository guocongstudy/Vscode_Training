package main

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"testing"
)

/*
#结构体标签和反射
-json的标签解析json
-yaml的标签解析yaml
-xorm gorm的标签标识db字段
-自定义标签
-原理是t.Field.Tag.Lookup("标签名")
*/
type Person2 struct {
	Name string `json:"name" yaml:"yaml_name" mage:"name"`
	Age  int    `json:"age" yaml:"yaml_age" mage:"age"`
	City string `json:"city" yaml:"yaml_city" mage:"city"`
}

//JSON解析
func jsonWork() {
	//对象marshal或字符串
	p := Person2{
		Name: "小一",
		Age:  15,
		City: "beijing",
	}
	data, err := json.Marshal(p)
	if err != nil {
		log.Printf("[json.marshal.err][err:%v]", err)
		return
	}
	log.Printf("[person.marshal.res][res:%v]", string(data))
	//从字符串解析成结构体
	p2Str := `
{
  "name":"李逵",
  "age":28,
  "city":"山东",
}
`
	var p2 Person2
	err = json.Unmarshal([]byte(p2Str), &p2)
	if err != nil {
		log.Printf("[json.unmarshal.err][err:%v]", err)
		return
	}
	log.Printf("[person2.unmarshal.res][res:%v]", p2)
}

//yaml读取文件
func yamlWork() {
	filename := "a.yaml"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("[ioutil.ReadFile.err:打开文件失败][err:%v]", err)
		return
	}
	p := &Person2{}
	err = yaml.UnmarshalStrict(content, p)
	if err != nil {
		log.Printf("[yaml.UnmarshalStrict.err:反marshal失败][err:%v]", err)
		return
	}
	log.Printf("[yaml.UnmarshalStrict.res][res:%v]", p)

}

func TestLabel(t *testing.T) {
	jsonWork()
	yamlWork()
}
