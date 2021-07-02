package main

import "fmt"

/*
	布尔值可以和 &&（AND）和 ||（OR）操作符结合，并且有短路行为，如果运算符左边的值已经可以确定整个布尔表达式的值，那么运算符右边的值将不再被求值
	&&的优先级比||高（&& 对应逻辑乘法，|| 对应逻辑加法，乘法比加法优先级要高）
	布尔值并不会隐式转换为数字值 0 或 1，反之亦然，必须使用 if 语句显式的进行转换
	不允许将整型强制转换为布尔型
*/

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func itob(i int) bool {
	return i != 0
}

func main() {
	fmt.Println(btoi(true))
	fmt.Println(btoi(false))
	fmt.Println(itob(2134))
	fmt.Println(itob(0))
}
