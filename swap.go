package main

import "fmt"

func swap(a, b int) {
	a, b = b, a
}

func swap2(a, b *int) {
	t := *a
	*a = *b
	*b = t
}

func swap3(a, b *int) {
	a, b = b, a
}

func swap4(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	x, y := 1, 2
	swap(x, y)
	fmt.Printf("x: %d, y: %d\n", x, y)
	a, b := 3, 4
	swap2(&a, &b)
	fmt.Printf("a: %d, b: %d\n", a, b)
	c, d := 5, 6
	swap3(&c, &d)
	fmt.Printf("c: %d, d: %d\n", c, d)
	e, f := 7, 8
	swap4(&e, &f)
	fmt.Printf("e: %d, f: %d\n", e, f)

	xx := 2
	xx = 3
	fmt.Println(xx)
}
