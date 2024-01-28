package main

import "fmt"

func main() {
	// 1-定义，定义一个具有10个数字的数组
	// c语言定义 int nums[10] = {1,2,3,4}
	// go语言定义 nums := [10]int{1, 2, 3, 4}(常用)
	// 其他位置会给你自动编译为0
	nums := [10]int{1, 2, 3, 4}
	// 2-遍历方式一
	for i := 0; i < len(nums); i++ {
		fmt.Println("i: ", i, "j :", nums[i])
	}

	// 方式二： for range ===> python支持
	for key, value := range nums {
		fmt.Println("key: ", key, " value: ", value)
	}

	// 在go语言中，如果想忽略一个值，可以使用_替代
	for _, value := range nums {
		fmt.Println(" value: ", value)
	}

	// 不定长数组定义
	// 3-使用make进行创建数组

}
