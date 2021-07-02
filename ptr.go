package main

import "fmt"

var (
	cat int = 1
	str string = "banana"
	house string = "Malibu Point 10880, 90265"
)

func print2() {
	ptr := &house
	fmt.Printf("%T\n", ptr)
	fmt.Printf("%p\n", ptr)
	val := *ptr
	fmt.Printf("%T\n", val)
	fmt.Printf("%s\n", val)
}



type Stu struct {
	I  int    `json:"i"`
	St string `json:"st"`
}

func (s Stu) console() string {
	fmt.Println(s.St)
	s.I = 1234
	return s.St
}

func console(s Stu) string {
	fmt.Println(s.St)
	return s.St
}

var ss Stu = Stu{123, "000"}

var i int = 1

var li []int = []int {1, 2, 3}

func updateList(list []int) {
	list[0] = 11
}

func updateIndex(i int) {
	i = 111
}

func main() {
	updateList(li)
	fmt.Println(li[0])
	updateIndex(li[0])
	fmt.Println(li[0])


	fmt.Printf("%p, %p\n", &cat, &str)
	print2()
	(&Stu{1, "123"}).console()
	(Stu{1, "123"}).console()
	ss.console()
	fmt.Println(ss)
	console(ss)

	i1 := i
	i2 := &i
	fmt.Println(i1)
	fmt.Println(*i2)
	fmt.Println(i2)

	str := new(string)
	*str = "hello"
	fmt.Println(str)
	fmt.Println(*str)
}
