# 单元测试相关

## 基本命令相关

### 测试函数
* 测试包下全部测试用例

```
  go test -v
```

* 指定包下某个单侧用例

```
  go test -run=${匹配函数名} -v
```

* 跳过某些用例

```
  go test -short
```

```
  func TestTimeConsuming(t *testing.T) {
      if testing.Short() {
          t.Skip("short模式下会跳过该测试用例")
      }
  }
```

* 自动生成单测

```
  go get -u github.com/cweill/gotests/...
```

* 测试覆盖率
  * 查看测试覆盖率

  ```
    go test -cover
  ```

  * 覆盖率相关的记录信息输出到一个文件

  ```
    go test -cover -coverprofile=c.out
  ```  

  * 生成html报告信息

  ```
    go tool cover -html=c.out
  ```

* assert断言

```
  go get github.com/stretchr/testify
```
### 基准测试
```
  go test -bench={func name}
  go test -bench={func name} -benchmem
```
* 测试结果
  * 这个是使用vscode自带的测试，输出了基础和内存信息
  * 如果是go test -bench=Split，那么只有次数和耗时
  * 如果是go test -bench=Split -benchmem，那么会打出内存信息
```
  === RUN   BenchmarkSplit
  BenchmarkSplit
  BenchmarkSplit-10       11278839               101.1 ns/op           112 B/op            3 allocs/op
  testing: BenchmarkSplit-10 left GOMAXPROCS set to 8
  PASS
  ok      GoBasic/SingleTest/Basic        1.694s
```
* 性能比较函数
* 基准测试只能得到给定操作的绝对耗时，但是在很多性能问题是发生在两个不同操作之间的相对耗时，比如同一个函数处理1000个元素的耗时与处理1万甚至100万个元素的耗时的差别是多少？再或者对于同一个任务究竟使用哪种算法性能最佳？我们通常需要对两个不同算法的实现使用相同的输入来进行基准比较测试。
* 性能比较函数通常是一个带有参数的函数，被多个不同的Benchmark函数传入不同的值来调用。举个例子如下
```
  // go test -bench=Fib //匹配
  func benchmark(b *testing.B, size int){/* ... */}
  func Benchmark10(b *testing.B){ benchmark(b, 10) }
  func Benchmark100(b *testing.B){ benchmark(b, 100) }
  func Benchmark1000(b *testing.B){ benchmark(b, 1000) }
```
* 默认情况下，每个基准测试至少运行1秒。如果在Benchmark函数返回时没有到1秒，则b.N的值会按1,2,5,10,20,50，…增加，并且函数再次运行。
```
  BenchmarkFib40-10              3         335705153 ns/op               0 B/op          0 allocs/op
```
* 最终的BenchmarkFib40只运行了两次，每次运行的平均值只有不到一秒。像这种情况下我们应该可以使用-benchtime标志增加最小基准时间，以产生更准确的结果。
```
  go test -bench=Fib40 -benchtime=20s
```
* **b.ResetTimer**之前的处理不会放到执行时间里，也不会输出到报告中，所以可以在之前做一些不计划作为测试报告的操作
```
  func BenchmarkSplit(b *testing.B) {
	time.Sleep(5 * time.Second) // 假设需要做一些耗时的无关操作
	b.ResetTimer()              // 重置计时器
	for i := 0; i < b.N; i++ {
		Split("沙河有沙又有河", "沙")
	}
}
```
* 并行测试（不是并发）
* func (b *B) RunParallel(body func(*PB))会以并行的方式执行给定的基准测试。
RunParallel会创建出多个goroutine，并将b.N分配给这些goroutine执行， 其中goroutine数量的默认值为GOMAXPROCS。用户如果想要增加非CPU受限（non-CPU-bound）基准测试的并行性， 那么可以在RunParallel之前调用SetParallelism 。RunParallel通常会与-cpu标志一同使用
```
  func BenchmarkSplitParallel(b *testing.B) {
	// b.SetParallelism(1) // 设置使用的CPU数
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Split("沙河有沙又有河", "沙")
		}
	  })
  }
```
* 还可以通过在测试命令后添加-cpu参数如go test -bench=. -cpu 1来指定使用的CPU数量。
* 为什么并行测试的结果更快？
  * 因为基础性能测试通过for循环，分配给多个cpu，而并行测试是通过goroutine测试，充分利用了GMP调度能力,但是并行测试没办法说明函数的真正耗时，因为函数的执行时间=总耗时/执行次数，但是在并行的情况下公式不成立（固定耗时内，并行运行的次数会更多）
  
### Setup和TearDown
* 测试程序有时需要在测试之前进行额外的设置（setup）或在测试之后进行拆卸（teardown）。
* TestMain：通过在*_test.go文件中定义TestMain函数来可以在测试之前进行额外的设置（setup）或在测试之后进行拆卸（teardown）操作。
  * 如果测试文件包含函数:func TestMain(m *testing.M)那么生成的测试会先调用 TestMain(m)，然后再运行具体测试。TestMain运行在主goroutine中, 可以在调用 m.Run前后做任何设置（setup）和拆卸（teardown）。退出测试的时候应该使用m.Run的返回值作为参数调用os.Exit。
  * 有时候我们可能需要为每个测试集设置Setup与Teardown，也有可能需要为每个子测试设置Setup与Teardown,见函数setupTestCase和setupSubTest

### 示例函数
* 示例函数只要包含了// Output:也是可以通过go test运行的可执行测试
* 优先级很高，只要gotest就会运行，错了会直接报出