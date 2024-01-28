package main

func main() {
	// go语言也有指针
	// 结构体成员调用时 c语言：ptr->name go语言：ptr.name

	// go语言在使用指针时，会使用内部的垃圾回收机制（gc ： garbage collector）
	// go语言可以返回栈上的指针，程序在编译的时候就确定了变量的分配位置
	// 编译的时候，如果发现有必要的话，就将变量分配到堆上
	name := "lily"
	ptr := &name

	println("name", *ptr)
	println("name ptr", ptr)
	// 02 - 使用new关键字来定义
	name2ptr := new(string)
	*name2ptr = "Duke"
	println("name2: ", *name2ptr)
	println("name2 ptr: ", name2ptr)

	tmpString := testPtr()
	*tmpString = "沙沟"
	// 可以看出res是深拷贝 并且nil不是一种类型
	res := tmpString

	tmpString = nil
	if res == nil {
		println("res 是空， nil")
	} else {
		println("res 是非空")
	}

	try := true
	println("try: ", try)
}

// 定义一个函数，返回一个string类型的指针，go语言返回值写在参数列表后面
func testPtr() *string {
	city := "深圳"
	ptr := &city
	return ptr
}
