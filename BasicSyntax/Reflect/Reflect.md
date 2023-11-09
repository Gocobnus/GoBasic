# 反射

## 变量的内在机制
* 类型信息：预先定义好的元信息。
* 值信息：程序运行过程中可动态变化的。
  
 ## 反射的定义
 * 反射是指在程序运行期间对程序本身进行访问和修改的能力。
 * 程序在编译时，变量被转换为内存地址，变量名不会被编译器写入到可执行部分。在运行程序时，程序无法获取自身的信息。
 * 支持反射的语言可以在程序编译期间将变量的反射信息，**如字段名称、类型信息、结构体信息等整合到可执行文件中，并给程序提供接口访问反射信息**，这样就可以在程序运行期间获取类型的反射信息，并且有能力修改它们。
 *  空接口可以存储任意类型的变量，那我们如何知道这个空接口保存的数据是什么呢？ 反射就是在运行时动态的获取一个变量的类型信息和值信息。
  
## reflect包
* 在Go语言中反射的相关功能由内置的reflect包提供，任意接口值在反射中都可以理解为由**reflect.Type**和**reflect.Value**两部分组成，并且reflect包提供了**reflect.TypeOf和reflect.ValueOf**两个函数来获取任意对象的Value和Type。
  
### TypeOf()
* 使用reflect.TypeOf()函数可以获得任意值的类型对象（reflect.Type），程序通过类型对象可以访问任意值的类型信息
  
#### type name和type kind
在反射中关于类型还划分为两种：类型（Type）和种类（Kind）。因为在Go语言中我们可以使用type关键字构造很多自定义类型，而种类（Kind）就是指底层的类型，但在反射中，当需要区分指针、结构体等大品种的类型时，就会用到种类（Kind）
* Go语言的反射中像数组、切片、Map、指针等类型的变量，它们的.Name()都是返回空
  
### ValueOf()
* 通过反射获取值:通过switch获取
* 通过反射设置变量的值:想要在函数中通过反射修改变量的值，需要注意函数参数传递的是值拷贝，必须传递变量地址才能修改变量值。而反射中使用专有的Elem()方法来获取指针对应的值。
* Elem只能对指针或者接口生效
  
## IsNil&IsValid
* IsNil()常被用于判断指针是否为空；IsValid()常被用于判定返回值是否有效。'
* IsNil()报告v持有的值是否为nil。v持有的值的分类必须是通道、函数、接口、映射、指针、切片之一；否则IsNil函数会导致panic。
* IsValid()返回v是否持有一个值。如果v是Value零值会返回假，此时v除了IsValid、String、Kind之外的方法都会导致panic。
  
## 结构体反射

### 与结构体相关的方法
* Field(i int)StructField	根据索引，返回索引对应的结构体字段的信息。
* NumField()int	返回结构体成员字段数量。
* FieldByName(name string)(StructField, bool)	根据给定字符串返回字符串对应的结构体字段的信息。
* FieldByIndex(index []int)StructField	多层成员访问时，根据 []int 提供的每个结构体的字段索引，返回字段的信息。
* FieldByNameFunc(match func(string) bool)(StructField,bool)	根据传入的匹配函数匹配需要的字段。
* NumMethod()int	返回该类型的方法集中方法的数目
* Method(int)Method	返回该类型方法集中的第i个方法
* MethodByName(string)(Method, bool)	根据方法名返回该类型方法集中的方法
  
### valueof的FieldByName和typeof的FieldByName区别
* valueOf的结果是字段对应的value值
* typeOf的结果是字段本身的信息
