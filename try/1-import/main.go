package main

import (
	"fmt"
	// 从目录文件夹开始一直找到最高位 sub是文件名 同时也是包名
	"try/1-import/add"
)

func main() {
	a := 1
	b := 2
	ans1 := add.Add(a, b)
	fmt.Println(ans1)
}
