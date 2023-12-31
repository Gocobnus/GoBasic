# Interface相关

## 接口是什么
* 接口就是规定了一个需要实现的方法列表
* 在 Go 语言中一个类型只要实现了接口中规定的所有方法，那么我们就称它实现了这个接口

## 接收者类型和接口
* 值接收者实现接口，结构体本身或者指针可以赋值给接口变量
* 指针接收者实现接口，结构体本身是无法赋值给接口变量

## 检查是否实现接口
* 编译期间检查
```
    var _ Singer = &Bird{}
```

## 空接口
* 使用空接口实现可以接收任意类型的函数参数。
* 使用空接口实现可以保存任意值的字典。map[string]interface{}


## 接口值
* 由于接口类型的值可以是任意一个实现了该接口的类型值，所以接口值除了需要记录具体值之外，还需要记录这个值属于的类型。
* 也就是说接口值由“类型”和“值”组成，鉴于这两部分会根据存入值的不同而发生变化，我们称之为接口的动态类型和动态值。
* 当将一个类型a赋值给接口时，接口的动态类型就是a，接口的值就是a的拷贝
* 只有当类型和值都为nil的时候 接口== nil
* 只有当动态类型和值都相同的结构才相同（如果动态类型不支持比较，比如slice，那么比较时会panic）
* 不能对一个空接口值调用任何方法，否则会产生panic。

## 类型断言
* v, ok := x.(T)
* 通过switch进行选择判断
```
func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}
```