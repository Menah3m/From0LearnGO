package main

import "fmt"

// 切片 slice

func main() {
	// 切片的定义
	var a []string
	var b []int

	// 定义切片的同时进行初始化
	var c = []bool{true, false}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	
	// 基于数组得到切片
	num := [8]int{12, 13, 14, 15, 16, 19, 20, 23}
	s1 := num[3:7]

	fmt.Println(s1)
	fmt.Printf("s1的长度是%d,容量是%d\n", len(s1), cap(s1))
	fmt.Printf("%T\n", s1)

	// 切片再次切片
	s2 := s1[:len(s1)]
	s3 := s2[:3]

	fmt.Println(s2)
	fmt.Println(s3)

	// make函数构造切片  make([]数据类型，切片的长度，切片的容量)
	s4 := make([]int, 5, 10)

	fmt.Println(s4)
	fmt.Printf("%T\n", s4)

	// 通过内置的len()和cap()获取切片的长度和容量
	fmt.Println(len(s4))
	fmt.Println(cap(s4))

	// 切片之间不能进行比较，即不能使用'=='来判断全部元素是否相等
	//  唯一可以进行的是 和nil的比较   nil是数组，切片，map等的零值（未初始化时默认的值）
	var n []int
	var m = []int{} 
	l := make([]int, 0)

	fmt.Println(n, len(n), cap(n))
	fmt.Println(m, len(m), cap(m))
	fmt.Println(l, len(l), cap(l))
	if n == nil {
		fmt.Println("n==nil")
	}
	if m != nil {
		fmt.Println("m!=nil")
	}
	if l != nil {
		fmt.Println("l!=nil")
	}

	// 切片的赋值拷贝
	e1 := make([]int, 3) //声明一个长度和容量都为3的int类型切片[0,0,0]
	e2 := e1             // 把e1复制给e2
	e2[1] = 100          // 改变e2的内容

	fmt.Println(e2)      // e2 -> [0,100,0]
	fmt.Println(e1)      // e1 -> [0,100,0]  e1的值也发生了变化
						 //  原因：e1和e2使用的底层数组相同，
						 //        改变e2则会改变底层数组，因此e1跟着发生变化
	// 切片的遍历
	se := []int{1, 2, 3, 4, 5}

	// 1. for循环
	for i := 0; i < len(se); i++ {
		fmt.Println(i, se[i])
	}

	fmt.Println()
	
	// 2. for range循环
	for index, value := range se {
		fmt.Println(index, value)
	}


	// 切片的扩容  append
	// 切片要初始化之后才能使用
	var ad []int     //此时没有申请内存
	for i := 0;i < 10; i++ {
		ad = append(ad,i) //append函数在添加元素时，可能会使得底层数组进行扩容，因此需要将结果返回给一个新的变量
		fmt.Printf("%v len:%d cap:%d ptr:%p\n",ad , len(ad), cap(ad), ad)
	}
	
	// append()可以添加多个元素，或者其他的切片
	bc := []int{11, 2, 3, 56}
	ad = append(ad, 1, 3, 5, 9)
	ad = append(ad, bc...)  // 注意：被添加的切片后面要有'...'

	fmt.Println(ad)


	// 切片的复制 copy
	co := []int{1, 2, 3, 4, 5}
	bo := make([]int, 5, 5)
	copy(bo, co)
	do := bo
	bo[0] = 100

	fmt.Println(co)
	fmt.Println(bo)
	fmt.Println(do)

	// 切片的删除   append(s[:index],a[index+1:]...)

	cities := []string{"北京", "上海", "广州", "深圳"}
	cities = append(cities[0:2], cities[3:]...)
	fmt.Println(cities)

	// 练习题1
	e()


}

// 练习题1
func e() {
	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++{
		a = append(a,fmt.Sprintf("%v", i))
	}
	fmt.Println(a)
}