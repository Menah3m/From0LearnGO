package main


import (
	"fmt"
	"os"
)


// 学生信息管理系统

//功能：
//  1.添加学生信息
//  2.编辑学生信息
//  3.展示所有学生信息

func showMenu(){
	fmt.Println("欢迎来到学生信息管理系统")
	fmt.Println("1. 添加学生")
	fmt.Println("2. 编辑学生信息")
	fmt.Println("3. 展示学生信息")
	fmt.Println("4. 退出系统")
}

// 获取用户输入的学生信息
func getInput() *student{
	var (
		id int
		name string
		class string
	)
	fmt.Println("请按要求输入学生信息")
	fmt.Print("请输入学生的学号：")
	fmt.Scanf("%d\n", &id)
	fmt.Print("请输入学生的姓名：")
	fmt.Scanf("%s\n", &name)
	fmt.Print("请输入学生的班级：")
	fmt.Scanf("%s\n", &class)
	//拿到用户输入的学生信息
	stu := newStudent(id, name, class) //调用student的构造函数构造学生信息
	return stu
	

}

func main(){
	sm := newStudentMgr()
	for {
	// 1. 打印系统菜单
	showMenu()
	// 2. 等待用户选择要执行的选项
	var input int
	fmt.Println("请输入你要输入的序号：")
	fmt.Scanf("%d\n", &input)
	if input > 4{
		fmt.Println("请输入正确的序号（1~4）:")
		fmt.Scanf("%d\n", &input)
	} else if input < 1{
		fmt.Println("请输入正确的序号（1~4）:")
		fmt.Scanf("%d\n", &input)
	} else {
		fmt.Println("用户输入的是", input)
	}
	
	// 3. 执行用户选择的动作
	switch input {
	case 1:
	// 添加学生
		stu := getInput()
		sm.addStudent(stu)
	case 2:
	// 编辑学生信息
		stu := getInput()
		sm.modifyStudent(stu)
	case 3:
	// 展示学生信息
		sm.showStudent()
	case 4:
	// 退出系统
		os.Exit(0)
	}

	}
}