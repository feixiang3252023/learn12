package main

import "fmt"

func main() {
	a := "stored in a"
	b := "stored in b"
	fmt.Println("a - ", a)
	// b is not being used - invalid code
}

/*在 Go 语言中，如果声明了一个变量而未使用它，
会导致编译错误。在你提供的代码中，变量 b 被声明但未在程序中使用。
这会导致编译错误*/
