package main

import "fmt"

func main() {
	x := 13 % 3 //计算余数
	fmt.Println(x)
	if x == 1 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}
	//判断奇偶性，并输出

	for i := 1; i < 70; i++ {
		if i%2 == 1 {
			fmt.Println("Odd")
		} else {
			fmt.Println("Even")
		}
	}
	//循环输出1到69的奇偶性
}
