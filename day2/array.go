package main

import "fmt"

func main() {
	var arr1 [2]string //{"1","2"}
	arr1 = [2]string{"1", "2"}
	fmt.Println(arr1)

	var slice1 []string //[]type
	slice1 = []string{"a", "b", "c"}
	slice1 = append(slice1, "e")
	fmt.Println(slice1)

	var slice2 []string
	//用数组表示下面这句的意思就是，数组长度为10，已经将三个位置进行了初始化[0,0,0,...]
	slice2 = make([]string, 3, 10)
	slice2 = append(slice2, "f", "b", "c")
	fmt.Println(slice2)

	var slice3 []string
	slice2 = make([]string, 0, 2)
	slice2 = append(slice2, "a", "b", "c")
	fmt.Println(slice3)

	var slice4 []int
	slice4 = make([]int, 4, 8)
	//len计算的是初始化的长度
	fmt.Println(slice4, len(slice4))

	//切片的复制
	var slice5 []int
	slice5 = []int{1, 2, 3, 4} //[0<1>,1<2>,2<3>,3<4>]
	slice6 := slice5[1:3]
	fmt.Println(slice6)

	//切片
	fmt.Println(slice5[1])
	var slice7 []int
	//copy函数只适合于切片不适合用于map等其他类型，它的返回结果为一个int型值，表示copy的长度

	reset := copy(slice7, slice5)
	fmt.Println(slice7)
	fmt.Println(reset)

	slice8 := []int{1, 2, 3, 4, 5}
	slice9 := []int{9, 6, 7}
	//方法一：
	copy(slice9, slice8)
	//方法二：
	copy(slice8[3:], slice9)
	fmt.Println(slice9)
	fmt.Println(slice8)

}
