package main

import "fmt"

func main() {
	// 标签 LABEL1
	// goto LABEL1 ===> 不会保留i的状态
	// continue LABEL1 ===> continue会跳到指定的位置但是会记录之前的状态， i变成1
	// break LABEL1 ===> 直接跳出指定位置的循环
LABEL1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println("i: ", i, " j: ", j)
			if j == 3 {
				goto LABEL1
			}
		}
	}
}
