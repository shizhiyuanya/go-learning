package main

import "fmt"

func main() {
	// 切片： slice，他的底层也是数组，可以动态改变长度
	// 定义一个切片包含多个地名
	// 定长
	//names := [10]string{"北京", "上海", "广州", "深圳"}
	// 不定长
	names := []string{"北京", "上海", "广州", "深圳"}

	for i, v := range names {
		fmt.Println("i: ", i, " v: ", v)
	}

	// 1.追加数据
	names1 := append(names, "海南")
	fmt.Println("names: ", names)
	fmt.Println("names1: ", names1)

	fmt.Println("追加元素前, name的长度：", len(names), " 容量: ", cap(names))

	names = append(names, "海南")
	fmt.Println("names追加元素后赋值给自己", names)
	//分配了两倍的容量
	fmt.Println("追加元素后, name的长度：", len(names), " 容量: ", cap(names))
	// 2.对于一个切片，不仅有长度len的概念，还有一个容量的概念cap()

}
