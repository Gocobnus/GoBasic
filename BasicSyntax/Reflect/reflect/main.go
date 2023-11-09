package main

import (
	"fmt"
	"reflect"
)

func GetReflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type of input is :, %v, type name is : %v, type kind is : %v\n", v, v.Name(), v.Kind())
}

type MyInt int


func GetReflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is : %v\n", int64(v.Int()))
	case reflect.Float32:
		fmt.Printf("type is float32, value is : %v\n", float32(v.Float()))
	case reflect.Float64:
		fmt.Printf("type is float64, value is : %v\n", float64(v.Float()))
	case reflect.Bool:
		fmt.Printf("type is bool, value is : %v\n", v.Bool())
	case reflect.Struct:
		fmt.Printf("type is struct, value is : %v\n", v)
	default:
		fmt.Printf("unexpect type\n")
	}

}

func SetReflectValue(x interface{}) {
	// elem只能对指针或者接口生效
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
	

}
func main() {

	// typeof
	a := 3.14
	GetReflectType(a)
	b := struct {
		name string
		age  int
	}{
		name:"gocobnus",
		age: 26,
	}
	GetReflectType(b)

	var c MyInt = 10
	GetReflectType(c)

	d := make([]int, 10)
	GetReflectType(d)

	type book struct {
		name string
	}
	e := book{
		name:"test",
	}
	GetReflectType(e)


	// valueof
	a1 := 3.13131
	GetReflectValue(a1)
	var b1 int64 = 100
	GetReflectValue(b1)
	var c1 bool = false
	GetReflectValue(c1)
	GetReflectValue(e)
	SetReflectValue(&b1)
	fmt.Println(b1)

	// nil&valid
	// *int类型空指针
	var a2 *int
	fmt.Println("var a *int IsNil:", reflect.ValueOf(a2).IsNil())
	var a3 []int
	fmt.Println("var a *int IsNil:", reflect.ValueOf(a3).IsNil())
	// nil值
	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())
	// 实例化一个匿名结构体
	b2 := struct{
		abc string
	}{
		abc:"abc",
	}
	// 尝试从结构体中查找"abc"字段
	fmt.Println("不存在的结构体成员:", reflect.ValueOf(b2).FieldByName("abc").IsValid())
	// 尝试从结构体中查找"abc"方法
	fmt.Println("不存在的结构体方法:", reflect.ValueOf(b2).MethodByName("abc").IsValid())
	// map
	c2:= map[string]int{}
	// 尝试从map中查找一个不存在的键
	fmt.Println("map中不存在的键：", reflect.ValueOf(c2).MapIndex(reflect.ValueOf("娜扎")).IsValid())
}


