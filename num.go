package main

import (
	"fmt"
	"math"
)

var (
	a uint = 32
	b uint = 34
	c uint
	d uint
)

func main() {
	c = a - b
	d = a + b
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	var e uint = 0
	var f uint = 1
	var g uint
	g = e - f
	fmt.Println(g)
	fmt.Println(math.MaxFloat32)
	fmt.Printf("%.20f\n", math.Pi)
	num := complex(1,3)
	fmt.Println(num, real(num), imag(num))

	boolean := 1==1 && 1==2
	fmt.Println(boolean)
}
