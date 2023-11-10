# 泛型

## 什么是泛型
* 泛型允许程序员在强类型程序设计语言中编写代码时使用一些以后才指定的类型，在实例化时作为参数指明这些类型。ーー换句话说，在编写某些代码或数据结构时先不提供值的类型，而是之后再提供。
* 泛型为Go语言添加了三个新的重要特性:
    * 函数和类型的类型参数。
      * 向函数提供类型参数(代码中min函数中为int和float64)称为实例化，类型实例化分两步进行：
        * 首先，编译器在整个泛型函数或类型中将所有类型形参（type parameters）替换为它们各自的类型实参（type arguments）。
        * 其次，编译器验证每个类型参数是否满足相应的约束。
        ```
            fmin := min[float64] // 类型实例化，编译器生成T=float64的min函数
            m2 = fmin(1.2, 2.3)  // 1.2
        ``` 
      * 类型也可以使用类型参数列表  
      ```
            type Slice[T int | string] []T

            type Map[K int | string, V float32 | float64] map[K]V

            type Tree[T interface{}] struct {
                left, right *Tree[T]
                value       T
            }
      ``` 
    * 将接口类型定义为类型集，包括没有方法的类型。
    * 类型推断，它允许在调用函数时在许多情况下省略类型参数。