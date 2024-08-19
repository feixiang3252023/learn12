package main

import (
	"fmt"       // 用于格式化输入和输出
	"io/ioutil" //提供ReadAll函数来读取数据。
	"log"       // 用于记录日志和处理错误。
	"net/http"  //用于发起 HTTP 请求和处理 HTTP 响应。
)

func main() {
	res, err := http.Get("http://www.geekwiseacademy.com/")
	if err != nil { //检查是否发生了错误
		log.Fatal(err)
	}
	page, err := ioutil.ReadAll(res.Body)
	//函数读取响应的主体（即网页的内容），结果保存在 page 变量中。同样地，错误信息会保存在 err 变量中。
	res.Body.Close()
	if err != nil { //检查是否发生了错误
		log.Fatal(err)
	}
	fmt.Printf("%s", page)
}
