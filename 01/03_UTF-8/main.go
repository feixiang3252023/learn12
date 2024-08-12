package main

import "fmt"

func main() {
	for i := 60; i < 122; i++ { //for 循环从 60 开始，一直到 121。每次循环，i 的值会自增。
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
		//%q：以字符（字符串）形式打印 i 对应的 Unicode 字符（如果 i 的值是一个有效的字符码）
	}
}

/*作用：通过这段代码，用户可以看到每个数字在不同进制下的表示形式，以及它们对应的 Unicode 字符，
  这有助于理解数字、字符和不同进制之间的关系*/
