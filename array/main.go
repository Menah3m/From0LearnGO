package main

import "fmt"

// 数组相关内容

func main() {
	// 数组定义  var 数组名  [元素数量]元素数据类型
	var a [3]int 
	var b [4]int
	fmt.Println(a)
	fmt.Println(b)
	

	// 数组初始化 
	// 1. 定义时使用 初始值列表 来初始化
	var cityArray = [4]string{"北京", "上海", "广州", "深圳"} 
	fmt.Println(cityArray)
	fmt.Println(cityArray[0])
	fmt.Println(cityArray[2])

	// 2. 编译器推导数组长度
	var boolArray = [...]bool{true, false, true, false}
	fmt.Println(boolArray)
	fmt.Println(len(boolArray))

	// 3. 使用索引值方式初始化
	var langArray = [...]string{1:"Golang", 3:"Python", 7:"Java"}
	fmt.Println(langArray)
	fmt.Printf("type of langArray is: %T\n", langArray)


	// 数组的遍历
	
	// 1. for循环遍历
	for i := 0; i < len(cityArray); i++{
		fmt.Println(cityArray[i])
	}

	// 2. for range 遍历
	for _, value := range cityArray { //如果不需要用到索引值，则可以用匿名变量“_”来代替index变量
		fmt.Println(value)
	}


	// 二维数组
	// 1. 定义     注意：二维数组只有外层长度可以让编译器推导，内层不行
	citiesArray := [...][2]string{
		{"北京", "西安"},
		{"上海", "杭州"},
		{"重庆", "成都"},
		{"广州", "深圳"},
	}
	fmt.Println(citiesArray[2][1])

	// 2. 遍历
	for _, v1 := range citiesArray {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}


	// 数组是 值 类型 =>赋值和传参时会赋值整个数组，因此只改变副本的值，不改变本身的值
	x := [3]int{1, 2, 3}
	y := [][]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(x)
	fmt.Println(y)
	f1(x)
	f2(y)
	fmt.Println(x)
	fmt.Println(y)

	// 练习1： 求数组{1，3，5，7，8}的和
	var e1 = []int{1, 3, 5, 7, 8}
	arraySum(e1)

	// 练习2: 两数的和 leetcode.1
	var e2 = []int{2, 7, 11, 15}
	var e3 = []int{2, 5, 5, 8, 9}
	target1 := 9
	target2 := 10
	twoSum(e2, target1)
	twoSum(e3, target2)

	
}

func f1(a [3]int) {
	a[0] = 100
}

func f2(a [][]int) {
	a[0][0] = 100
}

// 函数arraySum用于求和
func arraySum(a []int) {
	sum := 0
	for i := 0; i < len(a); i++{
		sum += a[i]
	}
	fmt.Println(sum)
}

func twoSum(nums []int, target int) {
	for i := 0; i < len(nums); i++ {
		for j := i+1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				fmt.Println(i,j)
			}
		}
	}
}
