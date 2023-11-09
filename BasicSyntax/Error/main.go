package main

import (
	"GoBasic/BasicSyntax/Error/MyError"

	"fmt"
)

// 自定义error
func TestError() error {
	return MyError.NewError(400, "this is test error")
}



//github的error包


func main() {
	err := TestError()
	if err != nil {
		// 当我们使用fmt包打印错误时会自动调用 error 类型的 Error 方法
		// 也就是会打印出错误的描述信息。
		fmt.Println(err)
		opErr, _ := err.(*MyError.OpError)
		fmt.Println(opErr.GetCode(), opErr.GetMsg())
		
	}
}

