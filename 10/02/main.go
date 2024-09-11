package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ { //外层循环
		for j := 0; j < 5; j++ { //内层循环
			fmt.Println(i, " - ", j) //打印组合
		}
	}
}
