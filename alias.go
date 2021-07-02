package main

import "fmt"

// 类型别名可以解决重构中最麻烦的类型名变更问题
// 不能在一个非本地的类型上定义新方法，即不能为 不在一个包中的类型 定义方法。

type NewInt int

type IntAlias = int

func main() {
	var a NewInt
	fmt.Printf("a type is: %T\n", a) // a type is: main.NewInt
	var b IntAlias
	fmt.Printf("b type is: %T\n", b) // b type is: int
}
