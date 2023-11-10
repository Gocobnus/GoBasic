package main

import "fmt"

func reverseWithGenerics[T any](s []T) []T {
	l := len(s)
	r := make([]T, l)
	for i, e := range s {
		r[l-i-1] = e
	}
	return r
}


// 通过|进行类型约束
func min[T int | float64] (a, b T) T {
	if a <= b {
		return a
	}
	return b
}

func main() {
	fmt.Println(min[int](1,2))
}