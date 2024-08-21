package main

import "fmt"

const metersToYards float64 = 1.09361

//定义了一个常量 metersToYards，它表示米与码之间的转换比例（1 米约等于 1.09361 码）

func main() {
	var meters float64
	//声明一个名为 meters 的浮点数变量，来存储用户输入的米数
	fmt.Print("Enter meters swam: ")
	//打印提示信息，要求用户输入游泳的米数
	fmt.Scan(&meters)
	//从标准输入读取用户输入的值，并将其存储在 meters 变量中
	yards := meters * metersToYards
	//计算 yards 的值，通过乘以转换常量 metersToYards
	fmt.Println(meters, " meters is ", yards, " yards.")
	//将米数和转换后的码数一起输出。
}
