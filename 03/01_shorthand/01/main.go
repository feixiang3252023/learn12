package main //声明了程序的包名为 main

import "fmt" //引入标准库fmt,用于格式化输入输出

func main() {
	//定义了程序的入口函数 main，所有 Go 程序都从这个函数开始执行
	a := 10
	b := "golang"
	c := 4.17
	d := true
	e := "Hello"
	f := `Do you like my hat?`
	g := 'M'
	/*定义了多个变量
	a 是一个整数
	b 是一个字符串
	c 是一个浮点数
	d 是一个布尔值
	e 是一个字符串
	f 是一个原样字符串（使用反引号定义）
	g 是一个字符变量*/

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)
	fmt.Printf("%v \n", g)
	/*使用 fmt.Printf 函数打印变量的值，
	 %v 是格式化动词，用于以默认格式输出值。
	每个变量的值后面都有一个换行符 \n，
	确保每个值被打印在新的一行。*/
}

/*总结：这段代码展示了如何在 Go 程序中
定义不同类型的变量，然后使用 fmt 包
输出这些变量的值。程序将输出
整数、字符串、浮点数、布尔值、原样字符串和字符的
默认格式表示。*/
