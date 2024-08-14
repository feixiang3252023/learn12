package main

import (
	"fmt"

	"github.com/feixiang3252023/learn12/04/01_package-scope/02_visibility/vis"
	//vis：这是一个自定义的可见性包
)

func main() {
	fmt.Println(vis.MyName)
	//打印了 vis 包中的公共变量 MyName 的值。由于 MyName 首字母大写，意味着它是一个导出（公共）变量，能够被其他包访问
	vis.PrintVar()
	//调用了 vis 包中的 PrintVar() 函数
}
