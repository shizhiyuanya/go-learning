package main

import "fmt" // goland会帮我们自动导入程序中使用的包

func main() {
	// 变量定义: var
	// 常量定义： const

	// 01-先定义变量 再赋值
	var name string
	name = "duke"
	var age int
	age = 20
	fmt.Println("name: ", name)
	// printf才有格式输出 println没有
	fmt.Printf("name is %s, %d\n", name, age)

	// 02 定义时直接复制
	var gender = "man"
	fmt.Println("gender: ", gender)

	// 03 定义直接赋值，使用自动推导
	address := "北京"
	println("address", address)
	test(1, "sbsbsbs")

	// 平行赋值
	i, j := 10, 11
	println("i: ", i, " j: ", j)
	i, j = j, i
	println("i: ", i, " j: ", j)

}

func test(a int, b string) {
	println(a, b)
}
