package courseware

import (
	"fmt"
	"strings"
	"testing"
)

//写一个map函数，将所有的字符串转换成大写
func TestStringsRevise(t *testing.T) {

	list := []string{"abc", "def", "fqp"}
	out := MapStrToUpper(list, func(item string) string {
		return strings.ToUpper(item)
	})
	fmt.Println(out)
}

//"ABC","DWDD","AD","ITEM"
func MapStrToUpper(data []string, fn func(string) string) []string {
	newData := make([]string, 0, len(data))
	for _, v := range data {
		newData = append(newData, fn(v))
	}
	return newData
}
