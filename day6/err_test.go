package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
)

/*
错误处理
##复杂的错误类型
-提供了LinkError,PathError,SyscallError
-上述error都是实现了error接口的错误类型
##自定义error
-errors.New 单独的error
-自定义结构体，在原始错误的基础上，封装出自己的错误信息

##golang1.13中的Error Wrapping 是错误嵌套


*/
type MyError struct {
	err error
	msg string
}

func (e *MyError) Error() string {
	return e.err.Error() + e.msg
}

func yalidateArgs(name string) (ok bool, err error) {
	//判断字符串是否以“mysql”结尾
	if strings.HasSuffix(name, "mysql") {
		return true, nil
	} else {
		return false, errors.New("字符串不是以mysql结尾")
	}
}

func TestErr(t *testing.T) {
	fmt.Println(yalidateArgs("dwdwdmysql"))
	file, err := os.Stat("test.txt")
	if err != nil {
		switch err.(type) {
		case *os.PathError:
			log.Printf("PathError")
		case *os.LinkError:
			log.Printf("LinkError")
		case *os.SyscallError:
			log.Println("SyscallError")
		default:
			log.Println("未知错误")
		}

	} else {
		fmt.Println(file)
	}

	//创建原始的错误
	err = errors.New("原始的错误")
	newErr := MyError{
		err: err,
		msg: "是研发的错误",
	}
	fmt.Println(newErr)
	fmt.Println(err)

	e := errors.New("原始错误01")
	w := fmt.Errorf("wrap了一个新的错误:%w", e)
	fmt.Println(w)

}
