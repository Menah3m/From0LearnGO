package main

import (
	"fmt"
	"reflect"
)

// 结构体 反射


type student struct{
	Name string `json:"name" ini:"s_name"` 
	Score int `json:"score" ini:"s_score"`
}

func main(){
	stu1 := student{
		Name:"张三",
		Score:80,
	}

	// 通过反射去获取结构体中的所有字段
	t := reflect.TypeOf(stu1)
	fmt.Printf("name:%v  kind:%v\n", t.Name(), t.Kind())

	// 遍历结构体所有的字段
	for i:=0; i<t.NumField();i++{
		//根据结构体字段的索引去取字段
		fieldObj := t.Field(i)
		fmt.Printf("Name:%v type:%v tag:%v\n", fieldObj.Name, fieldObj.Type, fieldObj.Tag)
		fmt.Println(fieldObj.Tag.Get("json"), fieldObj.Tag.Get("ini"))
	}

	fieldObj2, ok := t.FieldByName("Score")
	if ok {
		fmt.Printf("Name:%v type:%v tag:%v\n", fieldObj2.Name, fieldObj2.Type, fieldObj2.Tag)
	}
}