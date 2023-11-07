package main

import (
	"fmt"
	"unsafe"
)

// 类型别名和自定义类型
type MyInt int     //自定义类型
type YourInt = int // 类型别名

type Person struct {
	Name string
	Age  int
	City string
}

func main() {
	// 空结构体不占内存
	var p struct{}
	fmt.Println(unsafe.Sizeof(p))
	fmt.Println("test")
}

// 结构体
// 值类型既可以调用值接收者的方法，也可以调用指针接收者的方法；
// 指针类型既可以调用指针接收者的方法，也可以调用值接收者的方法。
