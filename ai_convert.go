/*
Go语言中的 strconv 包为我们提供了字符串和基本数据类型之间的转换功能
strconv 包中常用的函数包括 Atoi()、Itia()、parse 系列函数、format 系列函数、append 系列函数等
*/
package main

import (
	"fmt"
	"strconv"
)

/*
func Itoa(i int) string
	int 转换为 string
func Atoi(s string) i
	string 转换为 int
 */

/*
func ParseBool(str string) (bool, error)
	string 转化为 bool
	"1", "t", "T", "true", "TRUE", "True" -> true, nil
	"0", "f", "F", "false", "FALSE", "False" -> false, nil

func ParseInt(s string, base int, bitSize int) (i int64, err error)
	string 转化为 int
	base 指定进制，取值范围是 2 到 36。如果 base 为 0，则会从字符串前置判断，“0x”是 16 进制，“0”是 8 进制，否则是 10 进制
	bitSize 指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64

func ParseUint(s string, base int, bitSize int) (n uint64, err error)
	string 转化为 uint
	类似ParseInt，但不接受负数作为参数

func ParseFloat(s string, bitSize int) (f float64, err error)
	string 转化为 float
	bitSize 指定了返回值的类型，32 表示 float32，64 表示 float64
 */

/*
func FormatBool(b bool) string
	bool 转化为 string
	"true"/"false"

func FormatInt(i int64, base int) string
	int  转化为 string
	base 必须在 2 到 36 之间
	使用小写字母“a”到“z”表示大于 10 的数字

func FormatUint(i uint64, base int) string
	uint 转化为 string
	类似FormatInt

func FormatFloat(f float64, fmt byte, prec, bitSize int) string
	float 转化为 string
	prec 控制精度（排除指数部分）：当参数 fmt 为“f”、“e”、“E”时，它表示小数点后的数字个数；当参数 fmt 为“g”、“G”时，它控制总的数字个数。如果 prec 为 -1，则代表使用最少数量的、但又必需的数字来表示 f。
	fmt 表示格式，可以设置为“f”表示 -ddd.dddd、“b”表示 -ddddp±ddd，指数为二进制、“e”表示 -d.dddde±dd 十进制指数、“E”表示 -d.ddddE±dd 十进制指数、“g”表示指数很大时用“e”格式，否则“f”格式、“G”表示指数很大时用“E”格式，否则“f”格式。
	bitSize 表示参数 f 的来源类型（32 表示 float32、64 表示 float64），会据此进行舍入。
*/

/*
Append 系列函数用于将指定类型转换成字符串后追加到一个切片中
Append 系列函数和 Format 系列函数的使用方法类似，只不过是将转换后的结果追加到一个切片中
 */

func main() {
	fmt.Println(strconv.Itoa(123)) // "123"
	fmt.Println(strconv.Atoi("123")) // 123, nil
	fmt.Println(strconv.Atoi("a123")) // 0, strconv.Atoi: parsing "a123": invalid syntax
	fmt.Println(strconv.ParseBool("1")) // true, nil
	fmt.Println(strconv.ParseBool("FAlse")) // false, strconv.ParseBool: parsing "FAlse": invalid syntax
	fmt.Println(strconv.ParseInt("", 10, 64)) // false,  strconv.ParseInt: parsing "": invalid syntax
	fmt.Println(strconv.FormatBool(false)) // "false"
	fmt.Println(strconv.FormatFloat(3.1415926, 'E', -1, 64)) // "3.1415926E+00"
	fmt.Printf("%T %s\n", strconv.AppendBool([] byte("append: "), true), strconv.AppendBool([] byte("append: "), true)) // []uint8 "append: true"
}
