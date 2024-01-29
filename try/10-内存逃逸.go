package main

import "fmt"

func main() {
	p1 := testPtr1()
	fmt.Println("p1: ", p1)
}

func testPtr1() *string {
	city := "深圳"
	ptr := &city
	return ptr
}
