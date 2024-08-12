package main

import "fmt"

func main() {

	var a int
	var b string
	var c float64
	var d bool

	/*声明了四个变量：
	a 是 int 类型的整数。
	b 是 string 类型的字符串。
	c 是 float64 类型的双精度浮点数。
	d 是 bool 类型的布尔值。*/

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)

	/*使用 fmt.Printf 函数打印每个变量的默认值：
	%v 是格式化动词，用于以默认格式输出变量的值。
	每个变量的值后面都有一个换行符 \n，
	确保每个值输出在新的一行*/
	fmt.Println()
	//打印一个空行，使输出格式更整洁。
}
