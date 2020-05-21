package main

import "fmt"

// 结构体的继承

type Animal struct {
	name	string
}

func (a *Animal) move(){
	fmt.Printf("%s会动~\n", a.name)

}

type Dog struct {
	Feet	int8
	*Animal
}

func (d *Dog) wang() {
	fmt.Printf("%s会叫汪汪汪~\n", d.name)
}

func main(){
	d := &Dog{
		Feet:4,
		Animal:&Animal{
			name:"乐乐",
		},
	}

	d.wang()
	d.move()
}