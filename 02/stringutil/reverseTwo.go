package stringutil

func reverseTwo(s string) string { //reverseTwo 函数接收一个字符串 s，返回反转后的字符串
	r := []rune(s) //将字符串转换为 rune 切片，以正确处理Unicode字符
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	//for 循环用于交换切片 r 中的字符，从前到后和从后到前进行交换，直到中间为止
	return string(r)
	//将反转后的 rune 切片转换回字符串并返回。
}

// this demonstrates how an unexported function
// can be used by an exported function in the same package
/*一个未导出（首字母小写）的函数
（如 reverseTwo）可以在同一个包内由一个已导出
（首字母大写）的函数（如 Reverse）使用。
  这展示了包内函数的封装和访问权限。*/
