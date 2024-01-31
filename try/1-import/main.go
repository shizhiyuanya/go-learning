package main

import (
	// 改名
	SUB "try/1-import/sub"

	// 从目录文件夹开始一直找到最高位 sub是文件名 同时也是包名
	// .表示当前文件可以直接使用 不需要在用前缀了 (不推荐)
	. "try/1-import/add"
)

func main() {
	a := 1
	b := 2
	// 包名.方法
	ans1 := Add(a, b)
	ans2 := SUB.Sub(20, 10)
	//fmt.Println(ans1, " ", ans2)
	println(ans1, " ", ans2)

}
