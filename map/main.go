package main

import (
	"fmt"
	"math/rand"
	"sort"
)


//map(映射)
func main() {
	// 光声明map类型，没有初始化，则是一个零值 nil
	var a map[string]int
	fmt.Println(a == nil)

	// map的初始化
	a = make(map[string]int, 8)
	fmt.Println(a == nil)

	// map中添加k-v对
	a["王二"] = 100
	a["刘三"] = 102
	fmt.Println(a)
	fmt.Printf("a:%#v\n", a)
	fmt.Printf("type：%T\n", a)

	// 声明map的同时完成初始化
	b := map[int]bool{1:true, 2:false, 3:true}
	fmt.Printf("%#v\n", b)
	fmt.Printf("Type: %T\n", b)

	// 在使用map前必须进行初始化
	// var c map[int]int
	// c[100] = 100  //c这个map没有初始化，不能进行操作（因为没有申请内存，没有位置）
	// c[200] = 200
	// fmt.Println(c)

	// 判断map中某个k-v对是否存在
	var scoreMap = make(map[string]int, 8)
	scoreMap["王二"] = 100
	scoreMap["刘三"] = 102

	v, ok := scoreMap["张三"]
	if ok {
		fmt.Println("张三在scoreMap中", v)
	} else {
		fmt.Println("查无此人")
	}

	// 遍历map
	// 1.for range  (map是无序的，因此遍历结果与添加的顺序无关)
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

	// 1-1 只遍历key
	for k := range scoreMap {
		fmt.Println(k)
	}

	// 1-2 只遍历value
	for _, v := range scoreMap {
		fmt.Println(v)
	}

	// 2. 按照某个固定的顺序遍历map
	var scoreBigmap = make(map[string]int, 100)

	// 2-1 添加50个k-v对
	for i := 0;i < 51; i++ {
		key := fmt.Sprintf("stu%02d", i)  //格式化得到key
		value := rand.Intn(100)
		scoreBigmap[key] = value
	}
	// 2-2 正常情况下的遍历
	for k, v := range scoreBigmap {
		fmt.Println(k, v)
	}
	// 2-3 按照key的大小排序进行遍历  
	keys := make([]string, 0, 100) //先取出所有的key
	for key := range scoreBigmap {
		keys = append(keys, key)
	}
	sort.Strings(keys) //对key进行排序
	for _, key := range keys{     //按照排序后的key对map进行遍历
		fmt.Println(key, scoreBigmap[key])
	}

	 


	// 删除k-v对   通过delete()
	delete(scoreMap, "刘三")
	fmt.Println(scoreMap)



}





































































