package main

import "fmt"

func main() {

	// 1- 定义
	name := "duke"

	//需要换行，原生输出字符串时，使用反引号``
	usage := `./a.out <option>
		--h help
		-a	xxxx`
	println("name: ", name)
	println("usage: ", usage)

	// 2.长度，访问
	// go可以用len访问
	l1 := len(name)
	println("l1: ", l1)

	// 不需要加括号
	for i := 0; i < len(name); i++ {
		fmt.Printf("i: %d, v: %c\n", i, name[i])
	}
	l2 := len(usage)
	println("l2: ", l2)
	for i := 0; i < len(usage); i++ {
		fmt.Printf("i: %d, v: %c\n", i, usage[i])
	}
	i, j := "hello", "world"

	// 3-拼接
	fmt.Println("i + j = ", i+j)

	// 使用const修饰为常量，不能修改
	const address = "beijing"
	println("address: ", address)
}
