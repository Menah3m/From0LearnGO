package main

import (
	"fmt"
	"reflect"
)

// reflect

// reflectType 通过反射获取类型
func reflectType(x interface{}){
	// 不知道调用这个函数时会传入的变量
	// 方式1 通过类型断言（只能通过猜测）
	// 方式2  反射reflect获取
	obj := reflect.TypeOf(x)
	fmt.Println(obj, obj.Kind()) //kind可以得到底层的类型（复合类型）
	fmt.Printf("%T\n",obj)
}

// reflectValue 通过反射获取值
func reflectValue(x interface{}){
	v := reflect.ValueOf(x)   //通过reflect.Valueof()得到的原始值类型为reflect.Value，并不是该值的原始类型
	k := v.Kind()    // 拿到值对应的类型种类
	fmt.Printf("%v, %T\n", v, v)
	// 如何得到一个该变量传入时候的原始类型
	switch k {
	case reflect.Float32:
		// 把反射取到的值转换成int32类型的变量
		ret := float32(v.Float())
		fmt.Printf("%v,%T\n", ret, ret)
	case reflect.Int32:
		ret := int32(v.Int())
		fmt.Printf("%v,%T\n", ret, ret)
	}
}

// reflectSetValue 通过反射设置值
func reflectSetValue(x interface{}){
	v := reflect.ValueOf(x)
	// Elem() 用来根据指针取值
	k := v.Elem().Kind()
	switch k {
	case reflect.Int32:
		v.Elem().SetInt(100)
	case reflect.Float32:
		v.Elem().SetFloat(1.1)
	}
}

// Cat 自定义数据类型
type Cat struct{}

// Dog 自定义数据类型
type Dog struct{}
 
func main(){
	// reflect.Typeof()
	var a float32 = 1.234
	reflectType(a)
	var b int8 = 12
	reflectType(b)

	// 结构体类型
	var c Cat
	reflectType(c)
	var d Dog
	reflectType(d)

	// slice
	var e []int
	reflectType(e)
	var f []string
	reflectType(f)

	// reflect.Valueof()
	var aa int32 = 100
	reflectValue(aa)

	var bb float32 = 1.32
	reflectValue(bb)

	// reflectSetValue
	var aaa int32 = 10
	reflectSetValue(&aaa)
	fmt.Println(aaa)

	var bbb float32 = 100.1
	reflectSetValue(&bbb)
	fmt.Println(bbb)

}