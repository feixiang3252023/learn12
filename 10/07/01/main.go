package main

import "fmt"

func main() {
	for i := 250; i <= 340; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}
	foo := "a"
	fmt.Println(foo)
	//打印字符串变量 foo 的值。
	fmt.Printf("%T \n", foo)
	//使用 fmt.Printf 打印 foo 的数据类型
}
