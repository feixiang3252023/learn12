package main

import "fmt"

func main() {
	for i := 50; i <= 140; i++ {
		fmt.Printf("%v - %v - %v \n", i, string(i), []byte(string(i)))
	} //使用 for 循环从 50 迭代到 140
}
