/*
切片（slice）是对数组的一个连续片段的引用，是一个引用类型
这个片段可以是整个数组，也可以是由起始和终止索引标识的一些项的子集，需要注意的是，终止索引标识的项不包括在切片内
Go语言中切片的内部结构包含地址、大小和容量
 */

package main

import "fmt"

/*
生成切片语法
从数组或切片生成新的切片
slice[开始位置 : 结束位置]
	slice：表示目标切片对象

直接声明新的切片
var name [] Type

使用函数构造切片
make([]Type, size, cap)
	虽然没有初始化元素，但是分配了内存
	Type 指切片的元素类型
	size 指为这个类型分配多少个元素
	cap  指为切片预分配的元素数量，仅用于提前分配空间，这个值设定后不影响 size

从数组或切片生成新的切片拥有如下特性：
	取出的元素数量为：结束位置 - 开始位置
	取出元素不包含结束位置对应的索引，切片最后一个元素使用 slice[len(slice)] 获取
	当缺省开始位置时，表示从连续区域开头到结束位置
	当缺省结束位置时，表示从开始位置到整个连续区域末尾
	两者同时缺省时，与切片本身等效
	两者同时为 0 时，等效于空切片，一般用于切片复位

内建函数 append() 可以为切片动态添加元素
添加元素时，如果空间不足以容纳足够多的元素，切片容量会扩大到原来的2倍
 */

const h int = 30

var height [h]int

func heightInit()  {
	for i, _ := range height {
		height[i] = i + 1
	}
}

func main() {
	heightInit()

	// slice[start:stop]
	fmt.Println(height[10:15]) // [11 12 13 14 15]
	fmt.Println(height[:15]) // [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15]
	fmt.Println(height[10:]) // [11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30]
	fmt.Println(height[0:0]) // []
	fmt.Println(height[:]) // [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30]
	fmt.Println(height[0:len(height)]) // [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30]

	// []Type
	var slice2 []int
	var slice3 []int = []int{}
	fmt.Printf("slice2: %v, slice3: %v\n", slice2, slice3) // slice2: [], slice3: []
	fmt.Printf("len(slice2): %d, len(slice3): %d\n", len(slice2), len(slice3)) // len(slice2): 0, len(slice3): 0
	fmt.Println(slice2 == nil) // true
	fmt.Println(slice3 == nil) // false

	// make
	slice4 := make([]int, 2, 2)
	slice5 := make([]int, 2, 10)
	fmt.Println(slice4 == nil) // false
	fmt.Println(len(slice5)) // 2
	fmt.Printf("%T\n", slice5) // []int

	// append
	var slice6 []int
	for i:=0;i<h;i++ {
		slice6 = append(slice6, i) // 添加到末尾\
		fmt.Printf("len: %d, cap: %d, pointer: %p\n", len(slice6), cap(slice6), slice6) // 扩容时cap=cap*2
	}
	for i:=0;i<h;i++ {
		slice6 = append([]int{-i}, slice6...) // 添加到开头
		fmt.Printf("len: %d, cap: %d, pointer: %p\n", len(slice6), cap(slice6), slice6) // 扩容时cap=cap+4
	}
}
