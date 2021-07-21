package main

import "fmt"

//scan的用法：
//func main() {
//	//如果用调用写的函数，需要把自己封装一个包，传入gitee然后再导入，这样可能才可以。
//	//pkg.Hello()
//	//pkg.ExampleSwitch()
//	var (
//		name string
//		age  int
//	)
//	//给的是变量的地址
//	//scan是以空格分隔的
//	fmt.Println("请输入姓名和年龄，用空格当分隔符：")
//	//fmt.Scan(&name, &age)
//	targe := "郭聪 23"
//	fmt.Sscan(targe, &name, &age)
//	//fmt.Scanf("%s:%d", &name, &age)
//	fmt.Println(name, age)
//}

//for循环的用法：
func main() {
	//写法一：
	//for i := 0; i <= 10; i++ {
	//	fmt.Println(i)
	//}
	//写法二：
	//i:=1
	//for i<=10{
	//	fmt.Println(i)
	//	i++
	//}
	//写法三：
	//i:=1
	//for{
	//	fmt.Println(i)
	//	if i==10{
	//		return
	//	}
	//	i++
	//}
	//举例：
	//a := "dafefefg"
	//这个v打印出来的是阿斯科码
	//for i, v := range a {
	//	fmt.Println(i, v)
	//	//range是值对象，所以下面这个示例是不会改地址的
	//	if v == 'g' {
	//		v = 'x'
	//		fmt.Println("change g")
	//	}
	//	fmt.Println(i, v)

	//99乘法表
	//for m := 1; m < 10; m++ {
	//	for n := 1; n <= m; n++ {
	//		fmt.Printf("%d*%d=%d", m, n, m*n)
	//	}
	//	fmt.Println()
	//}

	//求素数
	//isP := true
	//for m := 2; m <= 100; m++ {
	//	for n := 2; n <= (m / n); n++ {
	//		if m%2 == 0 {
	//			isP = false
	//		}
	//	}
	//	if isP {
	//		fmt.Println(m)
	//	}
	//}

	//break
	//for i := 1; i <= 100; i++ {
	//	for k := 1; k <= 10; k++ {
	//		fmt.Print(k)
	//		if k == 5 {
	//			continue
	//		}
	//	}
	//	fmt.Println()
	//}

	//标签,一般都是大写，用冒号。

	//LOOP:
	//	for i := 1; i <= 100; i++ {
	//		if i == 50 {
	//			i++
	//			goto LOOP
	//		}
	//		i++
	//		fmt.Printf("i 的值为：%d\n", i)
	//
	//	}

	//数据类型转换
	//a := MyCustom("m1")
	//b := "string b"
	//能够把b的类型转换成a
	//fmt.Println(reflect.TypeOf(a))
	//fmt.Println(reflect.TypeOf(b))
	//
	//b = string(a)
	//fmt.Println(a, b)

	//var a int8 =10
	//var b int64 =20000000
	//a = int8(b)
	//
	//fmt.Println(a,b)

	//a := Age(10)
	//var b int = 12
	//fmt.Println(a, b)

	//高精度转低精度会丢失精度
	//a := 3.14124453
	//var b int
	//b = int(a)
	//fmt.Println(b)

	//strconv
	//a := "3"
	//var b int
	//b, err := strconv.Atoi(a)
	//fmt.Println(b, err)

	//转换成布尔
	//var b bool
	//var err error
	//b,err =strconv.ParseBool("true")
	//fmt.Println(b,err)
	//var f64 float64
	//f64,err =strconv.ParseFloat("3.1232432",64)
	//fmt.Println(f64,err)

	//var (
	//	i   int64
	//	err error
	//)
	////2代表的是二进制
	//i, err = strconv.ParseInt("1111", 2, 64)
	//fmt.Println(i, err)

	//指针
	//var ip *int
	//var fp *float32
	//
	//fmt.Println(ip, &fp)
	var ip *int = new(int)
	//*ip是指针地址对应的值
	fmt.Println(&ip, ip, *ip)
}

//type MyCustom string
//type Age int
