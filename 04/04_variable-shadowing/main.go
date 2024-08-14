package main

import "fmt"

func max(x int) int {
	return 42 + x
}

//定义一个名为 max 的函数，接受一个整数参数 x，返回 42 + x 的结果。

func main() {
	max := max(7)
	fmt.Println(max) // max is now the result, not the function
}

//在 main 函数中调用 max(7)，并将返回值赋值给 max。此时，max 不再是函数，而是返回值 49（因为 42 + 7 = 49）。

// don't do this; bad coding practice to shadow variables

/*变量遮蔽：在 main 函数中使用 max 变量名来保存函数返回值，遮蔽了外层的 max 函数。这意味着 max 这个变量在 main 中表示的是整型结果，
而非函数。这可能导致代码的可读性和可维护性下降。
编码实践：尽量避免使用相同的标识符来表示不同的概念，
特别是在同一作用域内。建议使用不同的名称来保持代码清晰和易于理解，比如将函数命名为 getMax 或 calculateMax。*/
