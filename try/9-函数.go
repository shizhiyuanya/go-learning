package main

import "fmt"

// 1.函数返回值在参数列表之后
// 2.如果有多个返回值，需要使用圆括号包裹
func test(a int, b int, c string) (int, string, bool) {
	return a + b, c, true
}

func main() {
	v1, s1, _ := test(10, 20, "sbsbsb")
	fmt.Println("v1: ", v1, " s1: ", s1)
}
