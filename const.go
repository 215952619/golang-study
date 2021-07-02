package main

import "fmt"

type ChipType int

// 当做枚举类型使用
const (
	None ChipType = iota
	CPU
	GPU
)

func (c ChipType) String() string {
	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}
	return "N/A"
}

const (
	FlagNone = 1 << iota
	FlagRed
	FlagGreen
	FlagBlue
)

func main() {
	fmt.Printf("%s %d\n", CPU, CPU)
	fmt.Println(CPU.String())
	fmt.Printf("FlagNone: %d, FlagRed: %d, FlagGreen: %d, FlagBule: %d\n", FlagNone, FlagRed, FlagGreen, FlagBlue)
	fmt.Printf("FlagNone: %b, FlagRed: %b, FlagGreen: %b, FlagBule: %b\n", FlagNone, FlagRed, FlagGreen, FlagBlue)
}
