package main

import "fmt"

// panic 和 recover
func a() {
	fmt.Println("func a")
}

func b() {
	defer func(){
		err := recover()
		if err != nil {
			fmt.Println("func b error")
		}
	}()
	panic("panic in b")
}

func c() {
	fmt.Println("func c")
}

func main() {
	a()
	b()
	c()
}