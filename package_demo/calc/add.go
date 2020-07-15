package calc

import "fmt"
import "github.com/Menah3m/From0LearnGo/package_demo/snow"

// 标识符首字母大写表示对外可见

// Name 定义一个测试的全局变量
var Name = "张三"

// Add 定义一个计算两个int类型数据的函数
func Add(x, y int) int {
	snow.Snow()
	return x + y
}

// init 在包导入时自动执行，没有参数也没有返回值
func init(){
	fmt.Println("calc.init()")
	fmt.Println(Name)
	
}