package main

import (
	"fmt"
	"os"
)

// 从命令输入参数，在switch中进行处理

func main() {
	// c: argc, **argv
	// Go: os.Args ==> 可以直接获取命令输入, 是一个字符串切片 []string
	cmds := os.Args

	fmt.Println("cmsd: ", cmds, " cmds len : ", len(cmds))

	for key, cmd := range cmds {
		fmt.Println("key: ", key, " cmd ", cmd)
	}

	//switch expr {
	//// 默认加 break 想要向下穿透加 fallthrough
	//
	//}
	if len(cmds) < 2 {
		fmt.Println("请输如正确参数！")
	}
	switch cmds[1] {
	case "hello":
		fmt.Println("Hello")
		fallthrough
	case "world":
		fmt.Println("world")
		//	不符合任何结果的时候输出的东西是default 无法穿透
		fallthrough
	default:
		fmt.Println(" default called ")
	}

}
