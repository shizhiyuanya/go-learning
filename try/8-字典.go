package main

import "fmt"

func main() {
	// 1. 定义一个字典
	// 学生id => 学生名字 idNames
	var idNames map[int]string // 定义一个map, 此时这个map是不能直接赋值的, 他是空的

	// 2. 分配空间, 使用make, 可以不指定长度，但是建议指定长度，这样性能更高
	idScore := make(map[int]float64)
	idNames = make(map[int]string, 10)
	idNames[0] = "duke"
	idNames[1] = "lyly"

	// 3. 定义时直接分配空间 (最好)
	//idNames1 := make(map[int]string, 10)

	// 4. 遍历map
	for key, value := range idNames {
		fmt.Println("key: ", key, " value: ", value)
	}

	// 5. 如何确定一个key是否在map中
	// 在map中不存在访问越界的问题，它被认为所有的key都是有效的, 所以访问一个不存在的key不会崩溃
	// 返回这个类型的零值
	// 零值: bool => false 数字 => 0 字符串 => 空
	name9 := idNames[9]
	if name9 == "" {
		// 空字符串
		fmt.Println("name9", name9)
	}
	fmt.Println("idScore[100]: ", idScore[100])
	// 无法通过value判断一个key是否存在, 因此我们需要一个能够校验key是否存在的方式
	value, ok := idNames[1] // 如果id=1是存在的 那么value就是key=1对应的值，ok返回true
	if ok {
		fmt.Println("id=1这个key是存在的， value为: ", value)
	} else {
		fmt.Println(" 不存在 ")
	}
	// 6.删除map中的元素
	// 使用自有函数中的 delete 来删除指定的key
	delete(idNames, 1)    // 有效删除
	delete(idNames, 1000) // 无效删除 不会报错
	fmt.Println("idNames删除后", idNames)

	// 并发任务处理的时候，需要对map进行上锁 // TODO:

}
