package main

import "fmt"

// 方法的定义实例

// Person 是一个结构体
type Person struct {
	name string
	age  int8
}

// NewPerson 是一个Person类型的构造函数
func NewPerson(name string, age int8) *Person{
	return &Person{
		name: name,
		age: age,
	}

}

// Dream 是为Person类型定义的方法
func (p Person)Dream(){
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)

}

// SetAge 是一个修改年龄的方法(接收者类型为指针类型)
func (p *Person) SetAge(newAge int8){
	p.age = newAge
}
// SetAge2 是另一个修改年龄的方法(接收者类型为值类型)
func (p Person) SetAge2(newAge int8){
	p.age = newAge
}


func main(){
	p1 := NewPerson("张三",int8(18))
	(*p1).Dream()

	fmt.Println(p1.age)
	p1.SetAge(int8(20))
	fmt.Println(p1.age)
	p1.SetAge2(int8(15))
	fmt.Println(p1.age)
}