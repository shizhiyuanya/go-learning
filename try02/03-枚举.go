package main

import "fmt"

// 在go语言中没有枚举类型

const (
	MONDAY = iota
	TUESDAY
	WEDNESDAY
	THURSDAY
	FRIDAY
	SATURDAY
	SUNDAY
	M, N = iota, iota
)

func main() {
	fmt.Println(MONDAY)
	fmt.Println(TUESDAY)
	fmt.Println(WEDNESDAY)
}
