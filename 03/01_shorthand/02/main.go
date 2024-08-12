package main

import "fmt"

func main() {

	a := 10
	b := "golang"
	c := 4.17
	d := true
	e := "Hello"
	f := `Do you like my hat?`
	g := 'M'

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", e)
	fmt.Printf("%T \n", f)
	fmt.Printf("%T \n", g)
	/*%T 是 Go 语言中 fmt 包的格式化动词，
	用于打印变量的类型。当你在 fmt.Printf 函数中
	使用 %T 时，它会输出该变量的具体数据类型*/
}

//与前一个代码相似，但是这个打印的是每个变量的类型，而不是值
