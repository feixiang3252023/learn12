package main

import "fmt"

var a = "this is stored in the variable a" // package scope
//在包作用域中定义变量 a，并赋值为一个字符串
var b, c string = "stored in b", "stored in c" // package scope
//同时定义变量 b 和 c，其作用域也是包级的，并分别赋值。
var d string // package scope
//定义变量 d，但未赋值，默认为空字符串

func main() {

	d = "stored in d" // declaration above; assignment here; package scope
	//对此前定义的变量 d 进行赋值。
	var e = 42 // function scope - subsequent variables have func scope:
	//在函数作用域内定义变量 e，并赋值。
	f := 43
	//使用短变量声明定义变量 f，并赋值。
	g := "stored in g"
	//使用短变量声明定义字符串变量 g。
	h, i := "stored in h", "stored in i"
	//一次性定义变量 h 和 i 并赋值。
	j, k, l, m := 44.7, true, false, 'm' // single quotes
	//使用短变量声明同时定义并赋值变量 j, k, l, m，这里注意 'm' 是 rune 类型
	n := "n" // double quotes
	//使用短变量声明定义变量 n 并赋值
	o := `o` // back ticks
	//使用反引号定义变量 o，表示原始字符串
	fmt.Println("a - ", a)
	fmt.Println("b - ", b)
	fmt.Println("c - ", c)
	fmt.Println("d - ", d)
	fmt.Println("e - ", e)
	fmt.Println("f - ", f)
	fmt.Println("g - ", g)
	fmt.Println("h - ", h)
	fmt.Println("i - ", i)
	fmt.Println("j - ", j)
	fmt.Println("k - ", k)
	fmt.Println("l - ", l)
	fmt.Println("m - ", m)
	fmt.Println("n - ", n)
	fmt.Println("o - ", o)
	//次打印所有定义的变量及其值。最终的输出会显示各个变量的内容。
}
