package main

import "fmt"

type calculation func(int, int) int

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func cal(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}


func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}


func main() {

	var c calculation
	c = add
	fmt.Println(c(1, 2))
	fmt.Println(cal(1, 8, add))

	// 匿名函数
	// 将匿名函数保存到变量
	add := func(x, y int) int {
		return x + y
	}
	add(10, 40)

	// 匿名函数立即执行
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 90)

	// 闭包=函数+引用环境
	// 这里如果 f1 := adder 那这个时候f1是一个 func() func(int) int 相当于函数本身
	f1 := adder() // f1是func(int) int 相当于函数的返回值
	fmt.Println(f1(10))
	fmt.Println(f1(20))

	// 进阶
	// f1 f2共享同一个base
	f1, f2 := calc(10)
	fmt.Println(f1(1), f2(2)) //11 9
	fmt.Println(f1(3), f2(4)) //12 8
	fmt.Println(f1(5), f2(6)) //13 7
}
