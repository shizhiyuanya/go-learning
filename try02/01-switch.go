package main

import (
	"fmt"
	"os"
)

// 从命令输入参数，在switch中进行处理

func main() {
	// c: argc, **argv
	// Go: os.Args ==> 可以直接获取命令输入, 是一个字符串切片
	cmds := os.Args

	for key, cmd := range cmds {
		fmt.Println("key: ", key, " cmd ", cmd)
	}

	//switch expr {
	//// 默认加 break 想要向下穿透加 fallthrough
	//
	//}

}
