package main

import "fmt"

func main() {
	switch "Mhi" { //使用 switch 语句来检查字符串 "Mhi" 的值
	case "Daniel":
		fmt.Println("Wassup Daniel")
	case "Medhi":
		fmt.Println("Wassup Medhi")
	case "Jenny":
		fmt.Println("Wassup Jenny")
	default:
		fmt.Println("Have you no friends?")
		//如果没有匹配的情况，执行 default 语句，打印 "Have you no friends?"
	}
}
