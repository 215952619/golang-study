/*
一维数组声明语法：
var 数组变量名 [元素数量]Type
	元素数量：数组的元素数量，可以是一个表达式，但最终通过编译期计算的结果必须是整型数值，元素数量不能含有到运行时才能确认大小的数值
	Type：可以是任意基本类型，包括数组本身，类型为数组本身时，可以实现多维数组

多维数组声明语法
var array_name [size1][size2]...[sizen] array_type
	size1、size2 等等为数组每一维度的长度
 */
package main

import "fmt"

var arr1 [3]int // 声明时会讲每个元素初始化为元素类型对应的零值
var arr2 [3]int = [...]int{1,2,3} // 数组长度的位置出现“...”，表示数组的长度是根据初始化值的个数来计算
var arr3 [3]int = [3]int{1,2,3}
var arr4 [3]int = [3]int{1:2} // 声明数组时初始化数组中指定元素
var arr5 [2][2]int // 声明多维数组
var arr6 [2][2]int = [2][2]int{1: {1: 2}} // 声明多维数组并初始化指定的元素

func main() {
	for i, v := range arr1 {
		// 默认为对应类型的零值
		fmt.Printf("arr[%d]: %d\n", i, v)
	}

	fmt.Println("`[...] int{1,2,3}` length is: ", len([...] int{1,2,3})) // `[...] int{1,2,3}` length is:  3

	// 只有数组的长度和元素类型全部相同时才能被认定为相同的数组类型
	// 相同的数组类型的每一项元素都相等时才能认定为数组相同
	// 只能数组类型相同的元素才能进行比较
	fmt.Println("arr1 == arr2 ----- ", arr1 == arr2) // arr1 == arr2 -----  false
	fmt.Println("arr1 == arr3 ----- ", arr1 == arr3) // arr1 == arr3 -----  false
	fmt.Println("arr2 == arr3 ----- ", arr2 == arr3) // arr2 == arr3 -----  true

	fmt.Println(arr4) // [0 2 0]

	fmt.Println(arr5) // [[0 0] [0 0]]
	fmt.Println(arr6) // [[0 0] [0 2]]
	fmt.Println(arr6[1][1]) // 2
}
