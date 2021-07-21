package main

import (
	"Vscode_demo/day6/log"
	"Vscode_demo/day6/pk1"
	"fmt"
	"testing"
)

/*
包的根本概念
##包是什么
-go里面的包是用来组织源码，是多个源文件的集合

##包的声明定义
-在源文件中加上“package xxx” 就可以声明xxx的包
-新建一个pk1文件1文件夹，和内部的func1.go
-创建文件 src/pk1/pk2/func2.go
-创建文件 src/pk1/pk2/func3.go
-创建文件 src/index.go
-执行index.go
-关系说明，import导入的是路径，非包名
-包名目录不强制一致，但推荐一致
-代码中引用包的成员变量或者函数时，使用的包名不是目录
-在同一目录下，所有文件的名名
-文件名不限制，但不能有中文

##包名要求
-包名一般小写，使用一个简短且有意义的名字
-包名中可以含有-等特殊符号

##别名导入(为了重复的包)
-创建src/log/a.go
-方法的名字容易冲突

##使用下划线导入
-目的是：有时候并非真的需要这些包，仅仅希望的是他int()函数执行而已
-举例：mysql :github.com/go-sql-driver/mysql

##main包
-package main 下面可以有多个文件，但所有文件只能有一个main
-package main 每个go应用有一个名为main的包

##包管理之verder（用不着了）
-最开始的时候

##包管理之mod
-go modules 是goland 1.11新加的特性

##设置go mod 和 go proxy
-设置两个环境变量
-GO111MODULE=on
-GOPROXY=https://goproxy.io,direct

### GO111MODULE 有三个值：off,on和auto(默认值)
- =off 代表go命令将不会支持module功能，按照原来的那种verder目录
- =on go命令会使用module功能，一点也不会去GOPATH目录查找
- =auto，默认值，go命令行会根据当前目录来决定是否启用module功能
  -当前目录在GOPROTH/src之外的，而且该目录还有go.mod开启
  -处于GOPATH没有go.mod等于off
-如果不是用go mod,go get会去master分支拉代码
-使用go.mod 可以选tag

##使用go mod 管理新的项目
-创建一个新的项目，go mod init



##go mod 使用go mod 有以下命令
 -module 语句指定报的名字
 -require 语句代码指定依赖的模块
 -replace 可以替换
 -exclude

-go.sum 文件夹记录，dependency tree
-mod包缓存位置在 $GOPATH/pkg/mod/xxx
-可以修改，go.mod中的模块，再sync就可以

##github升级版本，1.0.2
1.本地新建一个分支， git checkout -b v1 (切换成v1分支)
2.git add
3.git commit -m "xxx"
4.git tag v1.0.2(给一个新的标签)
5.git push --tags orgin v1
--到远程仓库检查v1.0.2
6.本地调用方，修改go的调用信息 go.mod v1.0.1变成v1.0.2
7.调用方GetHostName就会变成v1.0.2版本

06 主版本号变更
-版本号，v1.0.2代表版本号，朱版本号可能会不兼容
-0次版本号，2代表修正版本号
-把common-tools中的GetLoaclIp由一个返回值修改为两个返回值
-主板本变更的时候 go get不会去拉了
-手动get， go get github.com/guocongstudy/common-tools/v2

*/

func TestPkg(t *testing.T) {
	fmt.Println(pk1.PK1Var)
	pk1.PK1Func1()
	log.LogPrint()
}
