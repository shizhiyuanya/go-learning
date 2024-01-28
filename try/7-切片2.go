package main

import "fmt"

func main() {
	names := []string{"北京", "上海", "广州", "深圳"}

	// 想基于names创建一个新的数组

	// 切片可以几乎一个数组 灵活的创建新的数组
	names2 := names[0:3] // 左闭右开
	fmt.Println("names2: ", names2)

	names2[2] = "Hello"
	// 都会变
	fmt.Println("修改了names[2]之后, names2:", names2)
	fmt.Println("修改了names[2]之后, names: ", names)

	// 1. 如果从0元素开始截取，那么冒号左边的数字可以省略
	names3 := names[:3]
	fmt.Println("names3: ", names3)

	// 2. 如果截取到数组最后一个元素，那么冒号右边的数字可以省略
	names4 := names[3:]
	fmt.Println("names4: ", names4)

	// 3. 如果从左至右全部使用，那么冒号左右两边的数字都可以省略
	names5 := names[:]
	fmt.Println("names5: ", names5)

	// 4. 也可以基于一个字符串进行切片截取 ： 取字符串的子串 helloworld
	sub1 := "helloworld"[5:]
	fmt.Println("sub1 : ", sub1)
	// 5. 可以在创建空切片的时候，明确指定切片的长度和容量
	str2 := make([]int, 10, 20) // 第三个参数不是必须的， 如果没填写默认和长度相同
	fmt.Println("str2: ", str2)
	fmt.Println("str2 len : ", len(str2), " cap: ", cap(str2))

	// 如果想让切片完全独立于原始数组 可以使用copy函数
	namesCopy := make([]string, len(names))
	// names是一个数组, copy函数接受的参数类型是切片, 所以需要使用[:]将数组变成切片
	copy(namesCopy, names[:])
	fmt.Println("namesCopy: ", namesCopy)
	fmt.Println("namesCopy len", len(namesCopy), " cap: ", cap(namesCopy))
	namesCopy[0] = "香港"
	fmt.Println("namesCopy: ", namesCopy)
	fmt.Println("names本身: ", names)
}
