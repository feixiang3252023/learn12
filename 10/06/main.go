package main

import "fmt"

func main() {
	i := 0
	for {
		i++
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
		if i >= 50 {
			break
		}
	}
}

/*该程序的主要功能是打印从1到50之间的所有奇数。
由于使用了continue语句，偶数不会被打印。
最终输出将是1、3、5、7、9，直到49*/
