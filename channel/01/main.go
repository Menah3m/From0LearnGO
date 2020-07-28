package main

import (
	"fmt"
)

func main(){
	var ch1 chan int
	ch <- 10
	fmt.Println(ch)
}