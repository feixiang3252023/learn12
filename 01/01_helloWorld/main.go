package main //声明当前文件属于main包

import "fmt" // 导入了"fmt"包

//fmt 包提供了格式化输入和输出的功能，例如 Println、Printf、Scanf 等函数。

func main() {
	fmt.Println("Hello world!") //调用了fmt包里的Println函数，打印括号里的字符串
}

/*拓展：
fmt 包支持多种格式化功能，允许你以不同的格式和类型输出数据。一些常用的函数包括：
fmt.Printf: 允许格式化输出。
fmt.Sprintf: 返回一个格式化的字符串。
fmt.Scanf: 从标准输入读取格式化数据
*/
