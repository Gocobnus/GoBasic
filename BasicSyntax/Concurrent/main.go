package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// 使用 goroutine 和 channel 实现一个计算int64随机数各位数和的程序，例如生成随机数61345，计算其每个位数上的数字之和为19。
// 开启一个 goroutine 循环生成int64类型的随机数，发送到jobChan
// 开启24个 goroutine 从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
// 主 goroutine 从resultChan取出结果并打印到终端输出
var wg sync.WaitGroup

func GenRankNumber(jobChan chan int64) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		jobChan <- rand.Int63()
	}
	close(jobChan)
}

func CalcSum(jobChan chan int64, resultChan chan int64) {
	defer wg.Done()
	for ele := range jobChan {
		var s int64
		for ele > 0 {
			s += ele % 10
			ele = ele / 10
		}
		resultChan <- s
	}
}
func main() {
	jobChan := make(chan int64, 100)
	resultChan := make(chan int64, 100)
	wg.Add(1)
	go GenRankNumber(jobChan)

	for i := 0; i < 24; i++ {
		wg.Add(1)
		go CalcSum(jobChan, resultChan)
	}
	wg.Wait()
	close(resultChan)
	count := 0
	for ele := range resultChan {
		count++
		fmt.Printf("epoch: %v, result : %v\n", count, ele)

	}
}

// 也可以这么写，但是不如for range简单
var once sync.Once
func CalcSum2(jobChan chan int64, resultChan chan int64) {
	for {
		num,ok:= <-jobChan
		if !ok{
			once.Do(func(){
				close(resultChan)
			})
			break
		}
		var t int64
		for num > 0 {
			left := num % 10
			num = num / 10
			t += left
		}
		resultChan <- t
	}

	wg.Done()
}