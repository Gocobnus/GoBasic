package main

import (
	"encoding/json"
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
}
