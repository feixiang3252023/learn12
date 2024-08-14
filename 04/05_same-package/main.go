package main

import "fmt"

func main() {
	fmt.Println(x) //试图打印未声明的变量x
}

// IMPORTANT
// to run this code:
// go run *.go
// ---- OR ----
// go build
//./05_same-package
//这些指令用于在操作系统的命令行界面中执行 Go 程序。go run *.go 命令会编译并运行当前目录下的所有 Go 文件，但在本例中，由于 x 未被声明，程序依旧不会成功运行。
