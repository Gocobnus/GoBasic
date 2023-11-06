package main

import "fmt"

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}


func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}


func main() {
	// 1
	fmt.Println(f1()) //5
	fmt.Println(f2()) //6 x被赋值了5 之后x++
	fmt.Println(f3()) //5 x++ 与y无关
	fmt.Println(f4()) //5 defer是对x的副本操作

	// 2
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y)) // calc(aa, 1, 3) print:a, 1,2,3
	x = 10
	defer calc("BB", x, calc("B", x, y)) // calc(bb, 10, 12) print: b, 10,2, 12
	y = 20
	// print bb, 10,12, 22
	// print aa, 1,3,4

}
