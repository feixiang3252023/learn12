package main

import "fmt"

func main() {
	name := `Todd` // back-ticks work like double-quotes
	//短变量声明 := 声明变量 name，并将其初始化为原始字符串 "Todd"
	fmt.Println("Hello ", name)
}

/*使用反引号（back-ticks）来声明一个字符串。
反引号在 Go 中用于创建原始字符串字面量，
这意味着它们可以保留换行和特殊字符
，而不需要进行转义。*/
