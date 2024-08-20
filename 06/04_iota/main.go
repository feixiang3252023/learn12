package main

import "fmt"

const ( //在 const 块中，定义了一组常量
	a = iota // 0
	b        // 1
	c        // 2
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
