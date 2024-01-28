// 每个go程序 都必须有一个package main
// 每个go程序，都是.go结尾的
// go程序没有.h，没有.o,只有.go
// 一个package 包名 相当于命名空间
package main

// 导入一个标准包fmt
import "fmt"

func main() {
	// go语言不需要使用分号结尾
	fmt.Println("Hello World")
}
