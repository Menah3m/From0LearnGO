##### 数组定义
var 数组变量名 [元素数量]元素数据类型<br>

```go
 var a [3]int
 var b [4]int 
 a = b //不能这样做，因为此时a和b是不同的数据类型（长度不同）
```
<br>
数组是 值 类型 =>赋值和传参时会赋值整个数组，因此只改变副本的值，不改变本身的值<br>
<br>
数组的长度是固定的，并且是作为类型的一部分<br>

##### 数组初始化 
1. 定义时使用 初始值列表 来初始化<br>
```go
var cityArray = [4]string{"北京", "上海", "广州", "深圳"} 
```
2. 编译器推导数组长度<br>
```go
var boolArray = [...]bool{true, false, true, false}
```
3. 使用索引值方式初始化<br>
```go
var langArray = [...]string{1:"Golang", 3:"Python", 7:"Java"}
```

##### 数组的遍历
	
1. for循环遍历<br>

2. for range 遍历<br>

##### 二维数组
1. 定义 
```go
	citiesArray := [...][2]string{
		{"北京", "西安"},
		{"上海", "杭州"},
		{"重庆", "成都"},
		{"广州", "深圳"},
	}

```
注意：二维数组只有外层长度可以让编译器推导，内层不行<br>

2. 遍历<br>
```go
	for _, v1 := range citiesArray {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

```
