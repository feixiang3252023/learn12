package main

import (
	"fmt"
	github.com/feixiang32552023/learn12.git/02/stringutil
	//someAlias "github.com/GoesToEleven/GolangTraining/02_package/icomefromalaska"
	//someAlias是一种导入方式，它可以给导入的包起一个别名，使代码更加简洁
)

func main() { //func main() 定义了程序的主入口函数，执行程序的主要逻辑
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
	//调用 stringutil 包中的 Reverse 函数，反转字符串 !oG ,olleH 并打印结果。
	fmt.Println(stringutil.MyName)
	//打印 stringutil 包中的导出变量 MyName 的值。
	fmt.Println(winniepooh.BearName)
	//打印 winniepooh 包中定义的导出变量 BearName 的值。
}

/*总结：这段代码展示了如何在 Go 程序中导入并使用其他包，
  并调用这些包中的函数和变量。具体而言，
  它演示了字符串处理（反转字符串），
  以及访问不同包中的变量*/
